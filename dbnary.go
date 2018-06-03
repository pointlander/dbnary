// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate protoc --go_out=. dbnary.proto

package dbnary

import (
	"bytes"
	"io"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/pointlander/compress"
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
