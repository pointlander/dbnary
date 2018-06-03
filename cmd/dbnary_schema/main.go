// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"compress/bzip2"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/knakk/rdf"
)

// Suffix is a suffix
type Suffix struct {
	Name  string
	Index int
	Count int
}

// Prefix is a rdf prefix
type Prefix struct {
	Name, URI string
	Key       bool
	Count     int
	Suffixes  map[string]*Suffix
}

// Prefixes are a list of rdf prefixes for dbnary
var Prefixes = []Prefix{
	{
		Name: "olia",
		URI:  "http://purl.org/olia/olia.owl#",
	},
	{
		Name: "dbnary",
		URI:  "http://kaiko.getalp.org/dbnary#",
	},
	{
		Name: "vartrans",
		URI:  "http://www.w3.org/ns/lemon/vartrans#",
	},
	{
		Name: "lime",
		URI:  "http://www.w3.org/ns/lemon/lime#",
	},
	{
		Name: "skos",
		URI:  "http://www.w3.org/2004/02/skos/core#",
	},
	{
		Name: "synsem",
		URI:  "http://www.w3.org/ns/lemon/synsem#",
	},
	{
		Name: "rdfs",
		URI:  "http://www.w3.org/2000/01/rdf-schema#",
	},
	{
		Name: "lexinfo",
		URI:  "http://www.lexinfo.net/ontology/2.0/lexinfo#",
	},
	{
		Name: "wikt",
		URI:  "https://en.wiktionary.org/wiki/",
	},
	{
		Name: "decomp",
		URI:  "http://www.w3.org/ns/lemon/decomp#",
	},
	{
		Name: "rdf",
		URI:  "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
	},
	{
		Name: "dbetym",
		URI:  "http://etytree-virtuoso.wmflabs.org/dbnaryetymology#",
	},
	{
		Name: "lexvo",
		URI:  "http://lexvo.org/id/iso639-3/",
	},
	{
		Name: "dcterms",
		URI:  "http://purl.org/dc/terms/",
	},
	{
		Name: "ontolex",
		URI:  "http://www.w3.org/ns/lemon/ontolex#",
	},
	{
		Name: "xs",
		URI:  "http://www.w3.org/2001/XMLSchema#",
	},
	{
		Name: "eng",
		URI:  "http://kaiko.getalp.org/dbnary/eng/",
		Key:  true,
	},
}

const (
	eng = "http://kaiko.getalp.org/dbnary/eng/"
	// Page is a wiktionary page
	Page = "http://kaiko.getalp.org/dbnary#Page"
)

func main() {
	prefixes := make(map[string]*Prefix)
	for i := range Prefixes {
		prefix := &Prefixes[i]
		prefix.Suffixes = make(map[string]*Suffix)
		prefixes[prefix.URI] = prefix
	}
	add := func(term rdf.Term) {
		if term.Type() != rdf.TermIRI {
			return
		}
		iri := term.String()
		prefixIndex := strings.LastIndexAny(iri, "/#") + 1
		prefix := prefixes[iri[:prefixIndex]]
		if prefix == nil {
			panic(fmt.Errorf("%v", term))
		}
		prefix.Count++
		if prefix.Key {
			return
		}
		suffix := iri[prefixIndex:]
		s, ok := prefix.Suffixes[suffix]
		if ok {
			s.Count++
			return
		}
		code := len(prefix.Suffixes)
		prefix.Suffixes[suffix] = &Suffix{
			Name:  suffix,
			Index: code,
			Count: 1,
		}
	}
	printPrefixes := func(w io.Writer) {
		fmt.Fprintf(w, "// Copyright 2018 The dbnary Authors. All rights reserved.\n")
		fmt.Fprintf(w, "// Use of this source code is governed by a BSD-style\n")
		fmt.Fprintf(w, "// license that can be found in the LICENSE file.\n\n")
		fmt.Fprintf(w, "package dbnary\n")
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "// Prefixes are iri prefixes\n")
		fmt.Fprintf(w, "var Prefixes = []Prefix{\n")
		for i := range Prefixes {
			fmt.Fprintf(w, " {\n")
			fmt.Fprintf(w, "  Name: \"%v\",\n", Prefixes[i].Name)
			fmt.Fprintf(w, "  URI: \"%v\",\n", Prefixes[i].URI)
			if Prefixes[i].Key {
				fmt.Fprintf(w, "  Key: %v,\n", Prefixes[i].Key)
			}
			if len(Prefixes[i].Suffixes) > 0 {
				fmt.Fprintf(w, "  Suffixes: []string{\n")
				suffixes := make([]*Suffix, len(Prefixes[i].Suffixes))
				for _, value := range Prefixes[i].Suffixes {
					suffixes[value.Index] = value
				}
				sort.Slice(suffixes, func(i, j int) bool {
					return suffixes[i].Count > suffixes[j].Count
				})
				for j := range suffixes {
					fmt.Fprintf(w, "   \"%v\",\n", suffixes[j].Name)
				}
				fmt.Fprintf(w, "  },\n")
			}
			fmt.Fprintf(w, " },\n")
		}
		fmt.Fprintf(w, "}\n")
	}
	/*Prefixes[0].Suffixes["test"] = 0
	Prefixes[0].Suffixes["test1"] = 1
	Prefixes[0].Suffixes["test2"] = 2
	Prefixes[0].Suffixes["test3"] = 3
	printPrefixes(os.Stdout)*/

	input, err := os.Open("en_dbnary_ontolex.ttl.bz2")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	ttl := bzip2.NewReader(input)
	dec := rdf.NewTripleDecoder(ttl, rdf.Turtle)
	pages, definitions := 0, 0
	triple, err := dec.Decode()
	for err != io.EOF {
		add(triple.Subj)
		add(triple.Pred)
		add(triple.Obj)

		if strings.HasPrefix(triple.Subj.String(), "http://") {
			if !strings.HasPrefix(triple.Subj.String(), eng) {
				fmt.Println(triple)
			}
		} else {
			definitions++
		}
		if triple.Obj.String() == Page {
			pages++
		}
		triple, err = dec.Decode()
	}
	fmt.Println("pages=", pages)
	fmt.Println("definitions=", definitions)

	schema, err := os.Create("schema.go")
	if err != nil {
		panic(err)
	}
	defer schema.Close()
	sort.Slice(Prefixes, func(i, j int) bool {
		return Prefixes[i].Count > Prefixes[j].Count
	})
	printPrefixes(schema)
}
