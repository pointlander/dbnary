// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/knakk/rdf"
	"github.com/pointlander/dbnary"
)

// Suffix is a suffix
type Suffix struct {
	Name  string
	Count int
}

// Prefix is a rdf prefix
type Prefix struct {
	Name, URI    string
	Key          bool
	Count        int
	Suffixes     map[string]*Suffix
	SuffixesList []*Suffix
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
		Name: "bul",
		URI:  "http://kaiko.getalp.org/dbnary/bul/",
		Key:  true,
	},
	{
		Name: "deu",
		URI:  "http://kaiko.getalp.org/dbnary/deu/",
		Key:  true,
	},
	{
		Name: "ell",
		URI:  "http://kaiko.getalp.org/dbnary/ell/",
		Key:  true,
	},
	{
		Name: "eng",
		URI:  "http://kaiko.getalp.org/dbnary/eng/",
		Key:  true,
	},
	{
		Name: "spa",
		URI:  "http://kaiko.getalp.org/dbnary/spa/",
		Key:  true,
	},
	{
		Name: "fin",
		URI:  "http://kaiko.getalp.org/dbnary/fin/",
		Key:  true,
	},
	{
		Name: "fra",
		URI:  "http://kaiko.getalp.org/dbnary/fra/",
		Key:  true,
	},
	{
		Name: "ind",
		URI:  "http://kaiko.getalp.org/dbnary/ind/",
		Key:  true,
	},
	{
		Name: "ita",
		URI:  "http://kaiko.getalp.org/dbnary/ita/",
		Key:  true,
	},
	{
		Name: "jpn",
		URI:  "http://kaiko.getalp.org/dbnary/jpn/",
		Key:  true,
	},
	{
		Name: "lat",
		URI:  "http://kaiko.getalp.org/dbnary/lat/",
		Key:  true,
	},
	{
		Name: "lit",
		URI:  "http://kaiko.getalp.org/dbnary/lit/",
		Key:  true,
	},
	{
		Name: "mlg",
		URI:  "http://kaiko.getalp.org/dbnary/mlg/",
		Key:  true,
	},
	{
		Name: "nld",
		URI:  "http://kaiko.getalp.org/dbnary/nld/",
		Key:  true,
	},
	{
		Name: "nor",
		URI:  "http://kaiko.getalp.org/dbnary/nor/",
		Key:  true,
	},
	{
		Name: "pol",
		URI:  "http://kaiko.getalp.org/dbnary/pol/",
		Key:  true,
	},
	{
		Name: "por",
		URI:  "http://kaiko.getalp.org/dbnary/por/",
		Key:  true,
	},
	{
		Name: "rus",
		URI:  "http://kaiko.getalp.org/dbnary/rus/",
		Key:  true,
	},
	{
		Name: "hbs",
		URI:  "http://kaiko.getalp.org/dbnary/hbs/",
		Key:  true,
	},
	{
		Name: "swe",
		URI:  "http://kaiko.getalp.org/dbnary/swe/",
		Key:  true,
	},
	{
		Name: "tur",
		URI:  "http://kaiko.getalp.org/dbnary/tur/",
		Key:  true,
	},
}

var (
	// KeysByName are the prefixes that are primary keys by name
	KeysByName = make(map[string]Prefix)
	// KeysByURI are the prefixes that are primary keys by uri
	KeysByURI = make(map[string]Prefix)
	// PrefixesByName are the prefixes by name
	PrefixesByName = make(map[string]*Prefix)
)

func init() {
	for _, prefix := range Prefixes {
		if prefix.Key {
			if _, ok := KeysByName[prefix.Name]; ok {
				panic("duplicate keys")
			} else {
				KeysByName[prefix.Name] = prefix
			}

			if _, ok := KeysByURI[prefix.URI]; ok {
				panic("duplicate keys")
			} else {
				KeysByURI[prefix.URI] = prefix
			}
		}
	}

	for i := range Prefixes {
		prefix := &Prefixes[i]
		prefix.Suffixes = make(map[string]*Suffix)
		PrefixesByName[prefix.URI] = prefix
	}
}

const (
	file = "schema.go"
	// Page is a wiktionary page
	Page = "http://kaiko.getalp.org/dbnary#Page"
)

func generateSchema(file string) {
	add := func(term rdf.Term) {
		if term.Type() != rdf.TermIRI {
			return
		}
		iri := term.String()
		prefixIndex := strings.LastIndexAny(iri, "/#") + 1
		prefix := PrefixesByName[iri[:prefixIndex]]
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
		prefix.Suffixes[suffix] = &Suffix{
			Name:  suffix,
			Count: 1,
		}
	}

	input, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	ttl := bzip2.NewReader(input)
	dec := rdf.NewTripleDecoder(ttl, rdf.Turtle)
	pages, definitions := 0, 0
	triple, err := dec.Decode()
	keys := make(map[string]bool)
	for err != io.EOF {
		add(triple.Subj)
		add(triple.Pred)
		add(triple.Obj)

		if strings.HasPrefix(triple.Subj.String(), "http://") {
			s := triple.Subj.String()
			index := strings.LastIndexAny(s, "/#") + 1
			if _, ok := KeysByURI[s[:index]]; !ok {
				if !keys[s] {
					fmt.Println(triple)
					keys[s] = true
				}
			}
		} else {
			definitions++
		}
		if triple.Obj.String() == Page {
			pages++
		}
		triple, err = dec.Decode()
	}
	fmt.Println(file)
	fmt.Println("pages=", pages)
	fmt.Println("definitions=", definitions)
	fmt.Println("keys=", keys)
}

func printSchema(out *bytes.Buffer) {
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
				suffixes := make([]*Suffix, 0, len(Prefixes[i].Suffixes))
				for _, v := range Prefixes[i].Suffixes {
					suffixes = append(suffixes, v)
				}
				sort.Slice(suffixes, func(i, j int) bool {
					return suffixes[i].Count > suffixes[j].Count
				})
				for j := range suffixes {
					fmt.Fprintf(w, "   \"%v\",\n", suffixes[j].Name)
				}
				fmt.Fprintf(w, "  },\n")
				Prefixes[i].SuffixesList = suffixes
			}
			fmt.Fprintf(w, " },\n")
		}
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "\nconst (\n")
		for i := range Prefixes {
			fmt.Fprintf(w, "  ID_%v = %v\n", Prefixes[i].Name, i)
			suffixes := Prefixes[i].SuffixesList
			for j := range suffixes {
				fmt.Fprintf(w, "  ID_%v_%v = %v\n", Prefixes[i].Name, suffixes[j].Name, j)
			}
		}
		fmt.Fprintf(w, ")\n")
	}
	/*Prefixes[0].Suffixes["test"] = 0
	Prefixes[0].Suffixes["test1"] = 1
	Prefixes[0].Suffixes["test2"] = 2
	Prefixes[0].Suffixes["test3"] = 3
	printPrefixes(os.Stdout)*/

	sort.Slice(Prefixes, func(i, j int) bool {
		return Prefixes[i].Count > Prefixes[j].Count
	})
	printPrefixes(out)
}

func main() {
	for _, ttl := range dbnary.TTLFiles {
		generateSchema(ttl.Name)
	}
	var buffer bytes.Buffer
	printSchema(&buffer)

	out, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	fileSet := token.NewFileSet()
	code, err := parser.ParseFile(fileSet, file, &buffer, parser.ParseComments)
	if err != nil {
		buffer.WriteTo(out)
		panic(fmt.Errorf("%v: %v", file, err))
	}
	formatter := printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
	err = formatter.Fprint(out, fileSet, code)
	if err != nil {
		buffer.WriteTo(out)
		panic(fmt.Errorf("%v: %v", file, err))
	}
}
