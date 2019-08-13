// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"compress/bzip2"
	"flag"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/knakk/rdf"
	"github.com/pointlander/dbnary/utils"
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
	// TTLURL is the domain of dbnary
	TTLURL = "http://kaiko.getalp.org/static/ontolex/latest/"
)

// Download downloads the dbnary files
func Download() {
	for _, file := range utils.TTLFiles {
		head, err := http.Head(TTLURL + file.Name)
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
		stat, err := os.Stat("./" + file.Name)
		if err != nil || stat.ModTime().Before(last) || stat.Size() != int64(size) {
			fmt.Println("downloading", file, size, last)
			response, err := http.Get(TTLURL + file.Name)
			if err != nil {
				panic(err)
			}

			out, err := os.Create("./" + file.Name)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(out, response.Body)
			if err != nil {
				panic(err)
			}
			response.Body.Close()
			out.Close()

			err = os.Chtimes("./"+file.Name, last, last)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("skipping", file, size, last)
		}
	}
	return
}

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

const isoTemplate = `// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// http://www.iso639-3.sil.org/

package utils

// LanguageScope is the scope of the language
type LanguageScope uint8

const (
	// LanguageScopeIndividual is an individual language scope
	LanguageScopeIndividual LanguageScope = iota
	// LanguageScopeMacrolanguage is a macro language scope
	LanguageScopeMacrolanguage
	// LanguageScopeSpecial is a special language scope
	LanguageScopeSpecial
)

// LanguageType is a language type
type LanguageType uint8

const (
	// A is an ancient language type
	LanguageTypeAncient LanguageType = iota
	// LanguageTypeConstructed is a constructed language
	LanguageTypeConstructed
	// LanguageTypeExtinct is a extinct language
	LanguageTypeExtinct
	// LanguageTypeHistorical is a historical language
	LanguageTypeHistorical
	// LanguageTypeLiving is a living language
	LanguageTypeLiving
	// LanguageTypeSpecial is a special language
	LanguageTypeSpecial
)

// Language is a language
type Language struct {
	ID string
	Part2B string
	Part2T string
	Part1 string
	Scope LanguageScope
	Type LanguageType
	Name string
	Comment string
}

// Languages are a list of languages
var Languages = map[string]Language{
{{- range .}}
	"{{index . 0}}": {"{{index . 0}}", "{{index . 1}}", "{{index . 2}}", "{{index . 3}}", {{index . 4}}, {{index . 5}}, "{{index . 6}}", "{{index . 7}}"},
{{- end}}
}
`

// ISO generates iso-639-3 file
func ISO() {
	in, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer in.Close()
	out, err := os.Create("iso639-3.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	isoTemplate, err := template.New("iso").
		Parse(isoTemplate)
	if err != nil {
		panic(err)
	}
	languages := make(chan []string, 8)
	go func() {
		reader := bufio.NewReader(in)
		_, err = reader.ReadString('\n')
		for err == nil {
			var line string
			line, err = reader.ReadString('\n')
			parts := strings.Split(line, "\t")
			for i, part := range parts {
				parts[i] = strings.TrimSpace(part)
			}
			switch parts[4] {
			case "I":
				parts[4] = "LanguageScopeIndividual"
			case "M":
				parts[4] = "LanguageScopeMacrolanguage"
			case "S":
				parts[4] = "LanguageScopeSpecial"
			default:
				panic(fmt.Errorf("invalid scope %v", parts))
			}
			switch parts[5] {
			case "A":
				parts[5] = "LanguageTypeAncient"
			case "C":
				parts[5] = "LanguageTypeConstructed"
			case "E":
				parts[5] = "LanguageTypeExtinct"
			case "H":
				parts[5] = "LanguageTypeHistorical"
			case "L":
				parts[5] = "LanguageTypeLiving"
			case "S":
				parts[5] = "LanguageTypeSpecial"
			default:
				panic(fmt.Errorf("invalid type %v", parts))
			}
			languages <- parts
		}
		close(languages)
	}()
	err = isoTemplate.Execute(out, languages)
	if err != nil {
		return
	}

	/*fmt.Fprintf(out, "// Copyright 2018 The dbnary Authors. All rights reserved.\n")
	fmt.Fprintf(out, "// Use of this source code is governed by a BSD-style\n")
	fmt.Fprintf(out, "// license that can be found in the LICENSE file.\n\n")
	fmt.Fprintf(out, "// http://www.iso639-3.sil.org/\n\n")
	fmt.Fprintf(out, "package utils\n\n")
	fmt.Fprintf(out, "// Language is a language\n")
	fmt.Fprintf(out, "type Language struct {\n")
	fmt.Fprintf(out, "\tID string\n")
	fmt.Fprintf(out, "\tPart2B string\n")
	fmt.Fprintf(out, "\tPart2T string\n")
	fmt.Fprintf(out, "\tPart1 string\n")
	fmt.Fprintf(out, "\tScope string\n")
	fmt.Fprintf(out, "\tType string\n")
	fmt.Fprintf(out, "\tName string\n")
	fmt.Fprintf(out, "\tComment string\n")
	fmt.Fprintf(out, "}\n\n")
	fmt.Fprintf(out, "var Languages = map[string]Language{\n")
	reader, seen := bufio.NewReader(in), make(map[string]bool)
	_, err = reader.ReadString('\n')
	for err == nil {
		var line string
		line, err = reader.ReadString('\n')
		parts := strings.Split(line, "\t")
		if seen[parts[0]] {
			fmt.Printf(line)
			continue
		}
		for i, part := range parts {
			parts[i] = strings.TrimSpace(part)
		}
		fmt.Fprintf(out, "\t\"%s\": {\"%s\", \"%s\", \"%s\", \"%s\", \"%s\", \"%s\", \"%s\", \"%s\"},\n",
			parts[0], parts[0], parts[1], parts[2], parts[3], parts[4], parts[5], parts[6], parts[7])
		seen[parts[0]] = true
	}
	fmt.Fprintf(out, "}\n")*/
}

var (
	download = flag.Bool("download", false, "download the ttl files")
	schema   = flag.Bool("schema", false, "build the schema")
	iso      = flag.Bool("iso", false, "generate iso file")
	input    = flag.String("input", "", "the name of a file")
)

func main() {
	flag.Parse()

	if *iso {
		ISO()
		return
	}

	if *download {
		Download()
		return
	}

	if *schema {
		for _, ttl := range utils.TTLFiles {
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

		return
	}
}
