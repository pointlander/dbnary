// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"compress/bzip2"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/knakk/rdf"
	"github.com/pointlander/dbnary"
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

func build(db *dbnary.DB, file dbnary.TTLFile) {
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

var test = flag.Bool("test", false, "test mode")

func main() {
	flag.Parse()

	db := dbnary.OpenDB("dbnary.db", false)
	defer db.Close()
	for _, file := range dbnary.TTLFiles {
		build(db, file)
	}

	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	fmt.Println("sys=", stats.Sys)
}
