// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	"github.com/pointlander/dbnary"
)

const (
	// CacheSize the size of the write cache
	CacheSize = 1024
	eng       = "http://kaiko.getalp.org/dbnary/eng/"
	test      = false
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
		prefix, suffix := s[:index], s[index:]
		iri.Prefix = uint64(dbnary.PrefixesByURI[prefix].Index)
		if prefix == eng {
			iri.Key = suffix
		} else {
			iri.Suffix = uint64(dbnary.PrefixesByURI[prefix].SuffixesByName[suffix])
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

func build(db *dbnary.DB) {
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

	var entry *dbnary.Entry
	cache := make(map[string]*dbnary.Entry)
	writeEntry := func() {
		key := getKey()
		if key == "" {
			fmt.Println("invalid key:", current)
			return
		}

		cache[key], entry = entry, nil
		if len(cache) < CacheSize {
			return
		}

		err1 := db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("eng"))
			for i, v := range cache {
				value, err1 := proto.Marshal(v)
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
		cache = make(map[string]*dbnary.Entry)
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

		var err1 error
		entry, err1 = db.GetEntry(key)
		if err1 != nil {
			panic(err)
		}

		if entry == nil {
			entry = &dbnary.Entry{}
		}
	}

	if test {
		writeEntry = func() {}
		getEntry = func() {}
	}

	translations := make(map[string][]string)
	triple, err := dec.Decode()
	for err != io.EOF {
		subj := triple.Subj.String()
		if current != subj {
			if current != "" {
				writeEntry()
			}

			current = subj
			getEntry()
		}

		if entry != nil {
			trpl := &dbnary.Triple{
				Predicate: getTerm(triple.Pred),
				Object:    getTerm(triple.Obj),
			}
			entry.Triples = append(entry.Triples, trpl)

			if trpl.Predicate.Match(dbnary.ID_dbnary, dbnary.ID_dbnary_isTranslationOf) {
				words := translations[trpl.Object.Key]
				words = append(words, getKey())
				translations[trpl.Object.Key] = words
			}
		}

		triple, err = dec.Decode()
	}
	writeEntry()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("eng_translation"))
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

func main() {
	db := dbnary.OpenDB("dbnary.db", false)
	defer db.Close()
	build(db)
}
