package main

import (
	"bytes"
	"compress/bzip2"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/knakk/rdf"
	"github.com/pointlander/compress"
)

// Prefix is a rdf prefix
type Prefix struct {
	Name, URI string
	Key       bool
	Suffixes  map[string]int
	Index     int
	SuffixMap []string
}

const (
	// CacheSize the size of the write cache
	CacheSize = 1024
	eng       = "http://kaiko.getalp.org/dbnary/eng/"
	press     = true
	test      = false
)

var (
	prefixesURI  = make(map[string]*Prefix)
	prefixesName = make(map[string]*Prefix)
)

func init() {
	for i := range Prefixes {
		prefix := &Prefixes[i]
		prefix.Index = i
		suffixes := prefix.Suffixes
		prefix.SuffixMap = make([]string, len(suffixes))
		for k, v := range suffixes {
			prefix.SuffixMap[v] = k
		}
		prefixesURI[prefix.URI] = prefix
		prefixesName[prefix.Name] = prefix
	}
}

// Compress compresses some data
func Compress(input []byte, output io.Writer) {
	data, channel := make([]byte, len(input)), make(chan []byte, 1)
	copy(data, input)
	channel <- data
	close(channel)
	compress.BijectiveBurrowsWheelerCoder(channel).MoveToFrontCoder().FilteredAdaptiveBitCoder().Code(output)
}

// Decompress decompresses some data
func Decompress(input io.Reader, output []byte) {
	channel := make(chan []byte, 1)
	channel <- output
	close(channel)
	compress.BijectiveBurrowsWheelerDecoder(channel).MoveToFrontDecoder().FilteredAdaptiveBitDecoder().Decode(input)
}

func getTerm(term rdf.Term) *Term {
	switch term.Type() {
	case rdf.TermBlank:
		blank := &Term{
			Type: Term_Blank,
			Key:  term.String(),
		}
		return blank
	case rdf.TermIRI:
		iri := &Term{
			Type: Term_IRI,
		}
		s := term.String()
		index := strings.LastIndexAny(s, "/#")
		if index == -1 {
			panic("invalid iri")
		}
		index++
		prefix, suffix := s[:index], s[index:]
		iri.Prefix = uint64(prefixesURI[prefix].Index)
		if prefix == eng {
			iri.Key = suffix
		} else {
			iri.Suffix = uint64(prefixesURI[prefix].Suffixes[suffix])
		}
		return iri
	case rdf.TermLiteral:
		literal := &Term{
			Type:    Term_Literal,
			Literal: term.String(),
		}
		return literal
	}

	return nil
}

func build(db *bolt.DB) {
	input, err := os.Open("en_dbnary_ontolex.ttl.bz2")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	ttl := bzip2.NewReader(input)
	dec := rdf.NewTripleDecoder(ttl, rdf.Turtle)

	current := ""
	getKey := func() string {
		if strings.HasPrefix(current, eng) {
			return strings.TrimPrefix(current, eng)
		}
		return current
	}

	var entry *Entry
	cache := make(map[string]*Entry)
	writeEntry := func() {
		key := getKey()
		if key == "" {
			fmt.Println("invalid key:", current)
			return
		}

		cache[key] = entry

		if len(cache) > CacheSize {
			err1 := db.Update(func(tx *bolt.Tx) error {
				for i, v := range cache {
					value, err1 := proto.Marshal(v)
					if err1 != nil {
						return err1
					}
					if press {
						output := &bytes.Buffer{}
						Compress(value, output)
						compressed := &Compressed{
							Size: uint64(len(value)),
							Data: output.Bytes(),
						}
						value, err1 = proto.Marshal(compressed)
						if err1 != nil {
							return err1
						}
					}

					bucket := tx.Bucket([]byte("eng"))
					err2 := bucket.Put([]byte(i), value)
					if err2 != nil {
						return err2
					}
				}

				return nil
			})
			if err1 != nil {
				panic(err1)
			}
			cache = make(map[string]*Entry)
		}
	}
	getEntry := func() {
		key := getKey()
		if key == "" {
			fmt.Println("invalid key:", current)
			return
		}

		cached := cache[key]
		if cached != nil {
			entry = cached
			return
		}

		err1 := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("eng"))

			value := bucket.Get([]byte(key))
			if len(value) > 0 {
				if press {
					compressed := &Compressed{}
					err1 := proto.Unmarshal(value, compressed)
					if err1 != nil {
						return err1
					}
					value = make([]byte, compressed.Size)
					Decompress(bytes.NewReader(compressed.Data), value)
				}
				err1 := proto.Unmarshal(value, entry)
				if err1 != nil {
					return err1
				}
			}
			return nil
		})
		if err1 != nil {
			panic(err)
		}
	}

	if test {
		writeEntry = func() {}
		getEntry = func() {}
	}

	triple, err := dec.Decode()
	for err != io.EOF {
		subj := triple.Subj.String()
		if current != subj {
			if current != "" {
				writeEntry()
			}

			current, entry = subj, &Entry{}
			getEntry()
		}

		trpl := &Triple{
			Predicate: getTerm(triple.Pred),
			Object:    getTerm(triple.Obj),
		}
		entry.Triples = append(entry.Triples, trpl)

		triple, err = dec.Decode()
	}

	writeEntry()
}

var (
	buildFlag = flag.Bool("build", false, "build the database")
	printFlag = flag.String("print", "", "print")
)

func main() {
	flag.Parse()

	db, err := bolt.Open("dbnary.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err1 := tx.CreateBucketIfNotExists([]byte("eng"))
		return err1
	})
	if err != nil {
		panic(err)
	}

	if *buildFlag {
		build(db)
	}

	if *printFlag != "" {
		var display func(key, spaces string, depth int)
		display = func(key, spaces string, depth int) {
			entry := &Entry{}
			db.View(func(tx *bolt.Tx) error {
				bucket := tx.Bucket([]byte("eng"))
				value := bucket.Get([]byte(key))
				if len(value) > 0 {
					if press {
						compressed := &Compressed{}
						err1 := proto.Unmarshal(value, compressed)
						if err1 != nil {
							return err1
						}
						value = make([]byte, compressed.Size)
						Decompress(bytes.NewReader(compressed.Data), value)
					}
					err1 := proto.Unmarshal(value, entry)
					if err1 != nil {
						return err1
					}
				}
				return nil
			})
			printTerm := func(term *Term) string {
				link := ""
				switch term.Type {
				case Term_Blank:
					link = term.Key
					fmt.Print(term.Key)
				case Term_IRI:
					prefix := Prefixes[term.Prefix]
					fmt.Print(prefix.Name)
					fmt.Print(":")
					if prefix.Name == "eng" {
						link = term.Key
						fmt.Print(term.Key)
					} else {
						fmt.Print(prefix.SuffixMap[term.Suffix])
					}
				case Term_Literal:
					fmt.Print(term.Literal)
				}

				return link
			}
			for _, trpl := range entry.Triples {
				fmt.Print(spaces)
				printTerm(trpl.Predicate)
				fmt.Print(" ")
				link := printTerm(trpl.Object)
				fmt.Print("\n")
				if link != "" && depth < 0 {
					display(link, spaces+" ", depth+1)
				}
			}
		}
		display(*printFlag, "", 0)
	}

	/*total := 0
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("eng"))
		bucket.ForEach(func(k, v []byte) error {
			total += len(v)
			return nil
		})
		return nil
	})
	fmt.Println(total)*/
}
