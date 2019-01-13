// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate protoc --go_out=. dbnary.proto

package dbnary

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/knakk/rdf"
	"github.com/pointlander/compress"
	"github.com/pointlander/dbnary/utils"
)

const (
	// Press determines if the database is compressed
	Press = true
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

// GetKey gets the db key
func GetKey(triple rdf.Triple) (string, bool) {
	key := triple.Subj.String()
	index := strings.LastIndexAny(key, "/#")
	if index == -1 {
		return key, true
	}

	prefix := key[:index+1]
	if PrefixesByURI[prefix].Key {
		return strings.TrimPrefix(key, prefix), true
	}

	return "", false
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
			for _, file := range utils.TTLFiles {
				_, err1 := tx.CreateBucketIfNotExists([]byte(file.Key))
				if err1 != nil {
					return err1
				}
				_, err1 = tx.CreateBucketIfNotExists([]byte(fmt.Sprintf("%s_translation", file.Key)))
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
	entry = &Entry{}
	_, err = db.GetEntryForLanguage(key, "eng", entry)
	return
}

// GetEntryForLanguage looks up an entry for a given language
func (db *DB) GetEntryForLanguage(key, language string, entry *Entry) (valid bool, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(language))
		if bucket == nil {
			return fmt.Errorf("invalid language: %s", language)
		}
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
			err1 := proto.Unmarshal(value, entry)
			if err1 != nil {
				return err1
			}
			valid = true
		}
		return nil
	})
	return
}

// Mine is for mining the database
func (db *DB) Mine() {
	regexes := []*regexp.Regexp{
		regexp.MustCompile(`^\pN+$`),
		regexp.MustCompile(`^((\pN+([.,]\pN+)*)|([.,]\pN+)+)[.,]?$`),
		regexp.MustCompile(`^((\pN+\pL*)|(\pL+))$`),
		regexp.MustCompile(`^(((\pN+|\pL+)([.,]?(\pN+|\pL+))*)|([.,](\pN+|\pL+))+)[.,]?$`),
	}
	type Miss struct {
		failed  bool
		example string
		key     string
	}
	/*if regexC.Match([]byte("1.1")) {
		fmt.Println("matches")
	}*/
	for _, file := range utils.TTLFiles {
		fmt.Println(file.Key)
		misses := make([]Miss, len(regexes))
		err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(file.Key))
			bucket.ForEach(func(key []byte, value []byte) error {
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
					entry := &Entry{}
					err1 := proto.Unmarshal(value, entry)
					if err1 != nil {
						return err1
					}
					sense, isSense := "", false
					for _, triple := range entry.Triples {
						if triple.Predicate.Match(ID_dbnary, ID_dbnary_senseNumber) {
							sense = strings.TrimSpace(triple.Object.Literal)
						} else if triple.Predicate.Match(ID_rdf, ID_rdf_type) &&
							triple.Object.Match(ID_ontolex, ID_ontolex_LexicalSense) {
							isSense = true
						}
					}
					if isSense && sense != "" {
						for i, regex := range regexes {
							if !misses[i].failed && !regex.MatchString(sense) {
								misses[i].failed = true
								misses[i].example = sense
								misses[i].key = string(key)
							}
						}
					}
				}
				return nil
			})
			return nil
		})
		if err != nil {
			panic(err)
		}
		for i, miss := range misses {
			fmt.Println(i, !miss.failed, miss.example, miss.key)
		}
	}
}

// Check is for checking the database
func (db *DB) Check() {
	for _, file := range utils.TTLFiles {
		fmt.Println("checking:", file.Key)
		err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(file.Key))

			input, err := os.Open(file.Name)
			if err != nil {
				panic(err)
			}
			defer input.Close()
			ttl := bzip2.NewReader(input)
			dec := rdf.NewTripleDecoder(ttl, rdf.Turtle)

			triple, err := dec.Decode()
			for err != io.EOF {
				key, valid := GetKey(triple)
				if valid && key != "" {
					if bucket.Get([]byte(key)) == nil {
						fmt.Println("key not found:", key)
					}
				}
				triple, err = dec.Decode()
			}

			return nil
		})
		if err != nil {
			panic(err)
		}
	}
}

// PrintEntry prints the entry
func (db *DB) PrintEntry(key, spaces, lang string, max, depth int) {
	fmt.Println(key)
	var entry Entry
	valid, err := db.GetEntryForLanguage(key, lang, &entry)
	if err != nil {
		panic(err)
	}
	translation, err := db.GetTranslationForLanguage(key, lang)
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
			if prefix.Name == lang {
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
	if valid {
		for _, trpl := range entry.Triples {
			fmt.Print(spaces)
			printTerm(trpl.Predicate)
			fmt.Print(" ")
			link := printTerm(trpl.Object)
			fmt.Print("\n")
			if link != "" && depth < max {
				db.PrintEntry(link, spaces+" ", lang, max, depth+1)
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

// WriteNodes writes the nodes into lang
func (db *DB) WriteNodes(lang string, nodes *Node) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(lang))
		for nodes != nil {
			value, err := proto.Marshal(&nodes.Entry)
			if err != nil {
				return err
			}
			if Press {
				output := &bytes.Buffer{}
				Compress(value, output)
				compressed := &Compressed{
					Size: uint64(len(value)),
					Data: output.Bytes(),
				}
				value, err = proto.Marshal(compressed)
				if err != nil {
					return err
				}
			}

			err = bucket.Put([]byte(nodes.Key), value)
			if err != nil {
				return err
			}

			nodes = nodes.B
		}

		return nil
	})
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
	return db.LookupWordForLanguage(a, "eng")
}

// LookupWordForLanguage looks a word up in the dictionary for language
func (db *DB) LookupWordForLanguage(a, lang string) (word *Word, err error) {
	word = &Word{
		Word:      a,
		Relations: make(map[string][]string),
		Parts:     make([]*Part, 0),
	}
	getDefinition := func(a string) (definition string, err error) {
		var definitionEntry Entry
		valid, err := db.GetEntryForLanguage(a, lang, &definitionEntry)
		if err != nil || !valid {
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
		var entry Entry
		valid, err := db.GetEntryForLanguage(a, lang, &entry)
		if err != nil || !valid {
			return
		}
		partOfSpeech := -1
		parts := make(Definitions, 0, 8)
		getSense := func(a string) (err error) {
			var senseEntry Entry
			valid, err := db.GetEntryForLanguage(a, lang, &senseEntry)
			if err != nil || !valid {
				return
			}
			definition, sense := "", ""
			for _, triple := range senseEntry.Triples {
				if triple.Predicate.Match(ID_skos, ID_skos_definition) {
					definition, err = getDefinition(triple.Object.Key)
					if err != nil {
						return
					}
				} else if triple.Predicate.Match(ID_dbnary, ID_dbnary_senseNumber) {
					sense = triple.Object.Literal
				}
			}
			parts.Add(definition, sense)
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
		parts.Sort()
		part := &Part{
			Part:         PrefixesByName["lexinfo"].Suffixes[partOfSpeech],
			Translations: make(map[string][]string),
		}
		for _, p := range parts {
			part.Definitions = append(part.Definitions, p.Definition)
		}
		translations, err := db.GetTranslationForLanguage(a, lang)
		if err != nil {
			return
		}
		if translations != nil {
			for _, translation := range translations.Keys {
				var translationEntry Entry
				valid, err := db.GetEntryForLanguage(translation, lang, &translationEntry)
				if err != nil {
					return err
				}
				if !valid {
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

	var entry Entry
	key := strings.Replace(a, " ", "_", -1)
	valid, err := db.GetEntryForLanguage(key, lang, &entry)
	if err != nil || !valid {
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
	return db.GetTranslationForLanguage(key, "eng")
}

// GetTranslationForLanguage returns the translation for the given key and language
func (db *DB) GetTranslationForLanguage(key, language string) (translation *Translation, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(fmt.Sprintf("%s_translation", language)))
		if bucket == nil {
			return fmt.Errorf("no translations for %s", language)
		}
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
