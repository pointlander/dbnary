// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate protoc --go_out=. dbnary.proto

package dbnary

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/pointlander/compress"
)

const (
	// Press determines if the database is compressed
	Press = true
	// TTLURL is the domain of dbnary
	TTLURL = "http://kaiko.getalp.org/static/ontolex/latest/"
)

// Prefix is a rdf prefix
type Prefix struct {
	Name, URI      string
	Key            bool
	Suffixes       []string
	Index          int
	SuffixesByName map[string]int
}

var (
	// TTLFiles are the dbnary files
	TTLFiles = []string{
		"bg_dbnary_ontolex.ttl.bz2",
		"de_dbnary_ontolex.ttl.bz2",
		"el_dbnary_ontolex.ttl.bz2",
		"en_dbnary_ontolex.ttl.bz2",
		"es_dbnary_ontolex.ttl.bz2",
		"fi_dbnary_ontolex.ttl.bz2",
		"fr_dbnary_ontolex.ttl.bz2",
		"it_dbnary_ontolex.ttl.bz2",
		"id_dbnary_ontolex.ttl.bz2",
		"ja_dbnary_ontolex.ttl.bz2",
		"la_dbnary_ontolex.ttl.bz2",
		"lt_dbnary_ontolex.ttl.bz2",
		"mg_dbnary_ontolex.ttl.bz2",
		"nl_dbnary_ontolex.ttl.bz2",
		"no_dbnary_ontolex.ttl.bz2",
		"pl_dbnary_ontolex.ttl.bz2",
		"pt_dbnary_ontolex.ttl.bz2",
		"ru_dbnary_ontolex.ttl.bz2",
		"sh_dbnary_ontolex.ttl.bz2",
		"sv_dbnary_ontolex.ttl.bz2",
		"tr_dbnary_ontolex.ttl.bz2",
	}
	// PrefixesByURI are the prefixes indexed by URI
	PrefixesByURI = make(map[string]*Prefix)
	// PrefixesByName are the prefixes indexed by name
	PrefixesByName = make(map[string]*Prefix)
	// Relations are types of word relations
	Relations = map[uint64]bool{
		ID_dbnary_synonym:  true,
		ID_dbnary_hyponym:  true,
		ID_dbnary_antonym:  true,
		ID_dbnary_hypernym: true,
		ID_dbnary_meronym:  true,
		ID_dbnary_holonym:  true,
		ID_dbnary_troponym: true,
	}
)

func init() {
	for i := range Prefixes {
		prefix := &Prefixes[i]
		prefix.Index = i
		suffixes := prefix.Suffixes
		prefix.SuffixesByName = make(map[string]int, len(suffixes))
		for k, v := range suffixes {
			prefix.SuffixesByName[v] = k
		}
		PrefixesByURI[prefix.URI] = prefix
		PrefixesByName[prefix.Name] = prefix
	}
}

// Download downloads the dbnary files
func Download() {
	for _, file := range TTLFiles {
		head, err := http.Head(TTLURL + file)
		if err != nil {
			panic(err)
		}
		size, err := strconv.Atoi(head.Header.Get("Content-Length"))
		if err != nil {
			panic(err)
		}
		last, err := http.ParseTime(head.Header.Get("Last-Modified"))
		if err != nil {
			panic(err)
		}
		head.Body.Close()
		stat, err := os.Stat("./" + file)
		if err != nil || stat.ModTime().Before(last) || stat.Size() != int64(size) {
			fmt.Println("downloading", file, size, last)
			response, err := http.Get(TTLURL + file)
			if err != nil {
				panic(err)
			}

			out, err := os.Create("./" + file)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(out, response.Body)
			if err != nil {
				panic(err)
			}
			response.Body.Close()
			out.Close()

			err = os.Chtimes("./"+file, last, last)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("skipping", file, size, last)
		}
	}
	return
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

// Match returns true if prefix and suffix match
func (t *Term) Match(prefix, suffix uint64) bool {
	return t.Prefix == prefix && t.Suffix == suffix
}

// DB is a dbnary database
type DB struct {
	*bolt.DB
}

// OpenDB opens a dbnary db
func OpenDB(file string, readOnly bool) *DB {
	db, err := bolt.Open(file, 0600, &bolt.Options{ReadOnly: readOnly})
	if err != nil {
		panic(err)
	}

	if !readOnly {
		err = db.Update(func(tx *bolt.Tx) error {
			_, err1 := tx.CreateBucketIfNotExists([]byte("eng"))
			if err1 != nil {
				return err1
			}
			_, err1 = tx.CreateBucketIfNotExists([]byte("eng_translation"))
			return err1
		})
		if err != nil {
			panic(err)
		}
	}

	return &DB{
		DB: db,
	}
}

// Close closes the db
func (db *DB) Close() {
	db.DB.Close()
}

// GetEntry looks up an entry
func (db *DB) GetEntry(key string) (entry *Entry, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("eng"))
		value := bucket.Get([]byte(key))
		if len(value) > 0 {
			if Press {
				compressed := &Compressed{}
				err1 := proto.Unmarshal(value, compressed)
				if err1 != nil {
					return err1
				}
				value = make([]byte, compressed.Size)
				Decompress(bytes.NewReader(compressed.Data), value)
			} else {
				cp := make([]byte, len(value))
				copy(cp, value)
				value = cp
			}
			entry = &Entry{}
			err1 := proto.Unmarshal(value, entry)
			if err1 != nil {
				return err1
			}
		}
		return nil
	})
	return
}

// PrintEntry prints the entry
func (db *DB) PrintEntry(key, spaces string, max, depth int) {
	fmt.Println(key)
	entry, err := db.GetEntry(key)
	if err != nil {
		panic(err)
	}
	translation, err := db.GetTranslation(key)
	if err != nil {
		panic(err)
	}

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
				fmt.Print(prefix.Suffixes[term.Suffix])
			}
		case Term_Literal:
			fmt.Print(term.Literal)
		}

		return link
	}
	if entry != nil {
		for _, trpl := range entry.Triples {
			fmt.Print(spaces)
			printTerm(trpl.Predicate)
			fmt.Print(" ")
			link := printTerm(trpl.Object)
			fmt.Print("\n")
			if link != "" && depth < max {
				db.PrintEntry(link, spaces+" ", max, depth+1)
			}
		}
	}
	if translation != nil {
		for _, key := range translation.Keys {
			fmt.Print(spaces, " translation ")
			fmt.Println(key)
		}
	}
}

// Word a word
type Word struct {
	Word      string              `json:"word"`
	Relations map[string][]string `json:"relations"`
	Parts     []*Part             `json:"parts"`
}

// Part is a part of speech
type Part struct {
	Part         string              `json:"part"`
	Definitions  []string            `json:"definitions"`
	Translations map[string][]string `json:"translations"`
}

// LookupWord looks a word up in the dictionary
func (db *DB) LookupWord(a string) (word *Word, err error) {
	word = &Word{
		Word:      a,
		Relations: make(map[string][]string),
		Parts:     make([]*Part, 0),
	}
	getDefinition := func(a string) (definition string, err error) {
		definitionEntry, err := db.GetEntry(a)
		if err != nil || definitionEntry == nil {
			return
		}
		for _, triple := range definitionEntry.Triples {
			if triple.Predicate.Match(ID_rdf, ID_rdf_value) {
				definition = triple.Object.Literal
				return
			}
		}
		return
	}
	getPart := func(a string) (err error) {
		entry, err := db.GetEntry(a)
		if err != nil || entry == nil {
			return
		}
		partOfSpeech := -1
		type Sense struct {
			definition string
			sense      float64
		}
		var parts []Sense
		getSense := func(a string) (err error) {
			senseEntry, err := db.GetEntry(a)
			if err != nil || senseEntry == nil {
				return
			}
			definition, sense := "", 0.0
			for _, triple := range senseEntry.Triples {
				if triple.Predicate.Match(ID_skos, ID_skos_definition) {
					definition, err = getDefinition(triple.Object.Key)
					if err != nil {
						return
					}
				} else if triple.Predicate.Match(ID_dbnary, ID_dbnary_senseNumber) {
					sense, err = strconv.ParseFloat(triple.Object.Literal, 64)
					if err != nil {
						return
					}
				}
			}
			parts = append(parts, Sense{definition, sense})
			return
		}
		for _, triple := range entry.Triples {
			if triple.Predicate.Match(ID_lexinfo, ID_lexinfo_partOfSpeech) {
				partOfSpeech = int(triple.Object.Suffix)
			} else if triple.Predicate.Match(ID_ontolex, ID_ontolex_LexicalSense) ||
				triple.Predicate.Match(ID_ontolex, ID_ontolex_sense) {
				err = getSense(triple.Object.Key)
				if err != nil {
					return
				}
			}
		}
		sort.Slice(parts, func(i, j int) bool {
			return parts[i].sense < parts[j].sense
		})
		part := &Part{
			Part:         PrefixesByName["lexinfo"].Suffixes[partOfSpeech],
			Translations: make(map[string][]string),
		}
		for _, p := range parts {
			part.Definitions = append(part.Definitions, p.definition)
		}
		translations, err := db.GetTranslation(a)
		if err != nil {
			return
		}
		if translations != nil {
			for _, translation := range translations.Keys {
				translationEntry, err := db.GetEntry(translation)
				if err != nil {
					return err
				}
				if translationEntry == nil {
					continue
				}
				language, form := "", ""
				for _, triple := range translationEntry.Triples {
					if triple.Predicate.Match(ID_dbnary, ID_dbnary_targetLanguage) {
						language = PrefixesByName["lexvo"].Suffixes[triple.Object.Suffix]
					} else if triple.Predicate.Match(ID_dbnary, ID_dbnary_writtenForm) {
						form = triple.Object.Literal
					}
				}
				if language == "" || form == "" {
					continue
				}
				forms := part.Translations[language]
				forms = append(forms, form)
				part.Translations[language] = forms
			}
		}
		word.Parts = append(word.Parts, part)
		return
	}

	entry, err := db.GetEntry(a)
	if err != nil || entry == nil {
		return
	}
	for _, triple := range entry.Triples {
		if triple.Predicate.Match(ID_dbnary, ID_dbnary_describes) {
			err = getPart(triple.Object.Key)
			if err != nil {
				return
			}
		} else if triple.Predicate.Prefix == ID_dbnary &&
			Relations[triple.Predicate.Suffix] {
			relation := Prefixes[triple.Predicate.Prefix].Suffixes[triple.Predicate.Suffix]
			relations := word.Relations[relation]
			if relations == nil {
				relations = make([]string, 0, 1)
			}
			relations = append(relations, triple.Object.Key)
			word.Relations[relation] = relations
		}
	}
	return
}

// GetTranslation returns the translation for the given key
func (db *DB) GetTranslation(key string) (translation *Translation, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("eng_translation"))
		value := bucket.Get([]byte(key))
		if len(value) > 0 {
			if Press {
				compressed := &Compressed{}
				err1 := proto.Unmarshal(value, compressed)
				if err1 != nil {
					return err1
				}
				value = make([]byte, compressed.Size)
				Decompress(bytes.NewReader(compressed.Data), value)
			} else {
				cp := make([]byte, len(value))
				copy(cp, value)
				value = cp
			}
			translation = &Translation{}
			err1 := proto.Unmarshal(value, translation)
			if err1 != nil {
				return err1
			}
		}
		return nil
	})
	return
}
