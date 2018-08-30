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
	"strings"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/knakk/rdf"
	"github.com/pointlander/dbnary"
)

const (
	// CacheSize the size of the write cache
	CacheSize = 1024
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
	fmt.Println(file.Name, file.Key)
	var (
		primary     = []byte(file.Key)
		translation = []byte(fmt.Sprintf("%s_translation", file.Key))
	)
	input, err := os.Open(file.Name)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	ttl := bzip2.NewReader(input)
	dec := rdf.NewTripleDecoder(ttl, rdf.Turtle)

	current := ""
	getKey := func() (string, bool) {
		index := strings.LastIndexAny(current, "/#")
		if index == -1 {
			return current, true
		}

		prefix := current[:index+1]
		if dbnary.PrefixesByURI[prefix].Key {
			return strings.TrimPrefix(current, prefix), true
		}

		return "", false
	}

	var entry *dbnary.Entry
	cache := make(map[string]*dbnary.Entry)
	writeEntry := func() {
		key, valid := getKey()
		if !valid {
			return
		}
		if key == "" {
			fmt.Println("invalid key:", current)
			return
		}

		cache[key], entry = entry, nil
		if len(cache) < CacheSize {
			return
		}

		err1 := db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(primary)
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
		key, valid := getKey()
		if !valid {
			return
		}
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

	if *test {
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
				key, _ := getKey()
				words = append(words, key)
				translations[trpl.Object.Key] = words
			}
		}

		triple, err = dec.Decode()
	}
	writeEntry()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(translation)
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
}
