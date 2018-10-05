// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"compress/bzip2"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/knakk/rdf"
	"github.com/pointlander/dbnary"
	"github.com/pointlander/dbnary/utils"
)

func getTerm(term rdf.Term) *dbnary.Term {
	switch term.Type() {
	case rdf.TermBlank:
		blank := &dbnary.Term{
			Type: dbnary.Term_Blank,
			Key:  term.String(),
		}
		return blank
	case rdf.TermIRI:
		iri := &dbnary.Term{
			Type: dbnary.Term_IRI,
		}
		s := term.String()
		index := strings.LastIndexAny(s, "/#")
		if index == -1 {
			panic("invalid iri")
		}
		index++
		prefix, suffix := dbnary.PrefixesByURI[s[:index]], s[index:]
		iri.Prefix = uint64(prefix.Index)
		if prefix.Key {
			iri.Key = suffix
		} else {
			iri.Suffix = uint64(prefix.SuffixesByName[suffix])
		}
		return iri
	case rdf.TermLiteral:
		literal := &dbnary.Term{
			Type:    dbnary.Term_Literal,
			Literal: term.String(),
		}
		return literal
	}

	return nil
}

// Build builds the database
func Build(db *dbnary.DB, file utils.TTLFile) {
	lang := file.Key
	fmt.Println(file.Name, lang)
	input, err := os.Open(file.Name)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	ttl := bzip2.NewReader(input)
	dec := rdf.NewTripleDecoder(ttl, rdf.Turtle)

	translations, lru := make(map[string][]string), dbnary.NewLRU(11)
	triple, err := dec.Decode()
	for err != io.EOF {
		key, valid := dbnary.GetKey(triple)
		if valid && key != "" {
			node, found := lru.Get(key)
			if !found {
				_, err1 := db.GetEntryForLanguage(key, lang, &node.Entry)
				if err1 != nil {
					panic(err)
				}
			}

			trpl := &dbnary.Triple{
				Predicate: getTerm(triple.Pred),
				Object:    getTerm(triple.Obj),
			}
			node.Entry.Triples = append(node.Entry.Triples, trpl)
			if trpl.Predicate.Match(dbnary.ID_dbnary, dbnary.ID_dbnary_isTranslationOf) {
				words := translations[trpl.Object.Key]
				words = append(words, key)
				translations[trpl.Object.Key] = words
			}

			nodes := lru.Flush()
			if nodes != nil {
				db.WriteNodes(lang, nodes)
			}
		}
		triple, err = dec.Decode()
	}
	if lru.Head != nil {
		db.WriteNodes(lang, lru.Head)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(fmt.Sprintf("%s_translation", lang)))
		translation := dbnary.Translation{}
		for k, v := range translations {
			sort.Strings(v)
			translation.Keys = v
			value, err1 := proto.Marshal(&translation)
			if err1 != nil {
				return err1
			}
			if dbnary.Press {
				output := &bytes.Buffer{}
				dbnary.Compress(value, output)
				compressed := &dbnary.Compressed{
					Size: uint64(len(value)),
					Data: output.Bytes(),
				}
				value, err1 = proto.Marshal(compressed)
				if err1 != nil {
					return err1
				}
			}

			err2 := bucket.Put([]byte(k), value)
			if err2 != nil {
				return err2
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

var (
	print  = flag.String("print", "", "entry to print")
	depth  = flag.Int("depth", 0, "the depth to print entries")
	lookup = flag.String("lookup", "", "lookup a word")
	lang   = flag.String("lang", "eng", "language to use")
	mine   = flag.Bool("mine", false, "mine the database")
	check  = flag.Bool("check", false, "check the database")
	build  = flag.Bool("build", false, "build the database")
)

func main() {
	flag.Parse()

	if *build {
		db := dbnary.OpenDB("dbnary.db", false)
		defer db.Close()
		for _, file := range utils.TTLFiles {
			Build(db, file)
		}

		stats := runtime.MemStats{}
		runtime.ReadMemStats(&stats)
		fmt.Println("sys=", stats.Sys)
		return
	}

	db := dbnary.OpenDB("dbnary.db", true)
	defer db.Close()

	if *print != "" {
		db.PrintEntry(*print, "", *lang, *depth, 0)
		return
	}

	if *lookup != "" {
		word, err := db.LookupWordForLanguage(*lookup, *lang)
		if err != nil {
			panic(err)
		}
		data, err := json.MarshalIndent(word, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
		return
	}

	if *mine {
		db.Mine()
		return
	}

	if *check {
		db.Check()
		return
	}
}
