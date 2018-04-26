package main

import (
	"bytes"
	"compress/bzip2"
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
}

const (
	eng   = "http://kaiko.getalp.org/dbnary/eng/"
	press = true
)

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

func main() {
	prefixes, rprefixes := make(map[string]*Prefix), make(map[string]*Prefix)
	for i := range Prefixes {
		prefix := &Prefixes[i]
		prefix.Index = i
		prefix.Suffixes = make(map[string]int)
		prefixes[prefix.URI] = prefix
		rprefixes[prefix.Name] = prefix
	}

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
	getTerm := func(term rdf.Term) *Term {
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
			iri.Prefix = uint64(prefixes[prefix].Index)
			if prefix == eng {
				iri.Key = suffix
			} else {
				iri.Suffix = uint64(prefixes[prefix].Suffixes[suffix])
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
	writeEntry := func() {
		var value []byte
		value, err = proto.Marshal(entry)
		if err != nil {
			panic(err)
		}
		if press {
			output := &bytes.Buffer{}
			Compress(value, output)
			compressed := &Compressed{
				Size: uint64(len(value)),
				Data: output.Bytes(),
			}
			value, err = proto.Marshal(compressed)
			if err != nil {
				panic(err)
			}
		}

		err = db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("eng"))
			key := getKey()

			err1 := bucket.Put([]byte(key), value)
			if err1 != nil {
				return err1
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	i := 0
	triple, err := dec.Decode()
	for err != io.EOF && i < 1000 {
		subj := triple.Subj.String()
		if current != subj {
			if current != "" {
				writeEntry()
			}

			current, entry = subj, &Entry{}
			err = db.View(func(tx *bolt.Tx) error {
				bucket := tx.Bucket([]byte("eng"))
				key := getKey()
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
			if err != nil {
				panic(err)
			}
		}

		trpl := &Triple{
			Predicate: getTerm(triple.Pred),
			Object:    getTerm(triple.Obj),
		}
		entry.Triples = append(entry.Triples, trpl)

		i++
		triple, err = dec.Decode()
	}

	writeEntry()

	total := 0
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("eng"))
		bucket.ForEach(func(k, v []byte) error {
			total += len(v)
			return nil
		})
		return nil
	})
	fmt.Println(total)
}
