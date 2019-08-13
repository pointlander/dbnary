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
	"html/template"
	"io"
	"net/http"
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

// EntryTemplate is a entry page
const EntryTemplate = `<html>
 <head>
  <title>{{clean .Word}}</title>
	<style>
		.accordion {
			background-color: #eee;
			color: #444;
			cursor: pointer;
			padding: 18px;
			width: 50%;
			text-align: left;
			border: none;
			outline: none;
			transition: 0.4s;
		}
		.active, .accordion:hover {
			background-color: #ccc;
		}
		.accordion:after {
			content: '\02795';
			font-size: 13px;
			color: #777;
			float: right;
			margin-left: 5px;
		}
		.active:after {
			content: "\2796";
		}
		.panel {
			padding: 0 18px;
			background-color: white;
			overflow: hidden;
		}
	</style>
 </head>
 <body>
	<h1>{{clean .Word}}</h1>
	{{$wordLanguage := .Language}}
	<table>
		<tr>
{{range $relation, $words := .Relations}}
			<td style="vertical-align: top; padding: 20px;">
				<h2>{{$relation}}</h2>
				<ul>
{{range $words}}
					<li><a href="/{{$wordLanguage}}/{{.}}">{{clean .}}</a></li>
{{end}}
				</ul>
			</td>
{{end}}
		</tr>
  </table>
{{range .Parts}}
  <h2>{{.Part}}</h2>
	<ul>
{{range .Definitions}}
	 <li>{{.}}</li>
{{end}}
  </ul>
	<h3 class="accordion">Translations</h3>
	<div class="panel">
{{range $language, $translations := .Translations}}
		<h4>{{language $language}}</h4>
		<ul>
{{range $translations}}
{{if supports $language}}
			<li><a href="/{{$language}}/{{.}}">{{clean .}}</a></li>
{{else}}
			<li>{{clean .}}</li>
{{end}}
{{end}}
		</ul>
{{end}}
	</div>
{{end}}
	<script>
		var panels = document.getElementsByClassName("panel");
		for (i = 0; i < panels.length; i++) {
			panels[i].style.display = 'none';
		}
		var accordions = document.getElementsByClassName("accordion");
		var i;
		for (i = 0; i < accordions.length; i++) {
			accordions[i].addEventListener("click", function() {
				this.classList.toggle("active");
				var panel = this.nextElementSibling;
				if (panel.style.display === "block") {
					panel.style.display = "none";
				} else {
					panel.style.display = "block";
				}
			});
		}
	</script>
 </body>
</html>
`

// Dictionary a dictionary server
type Dictionary struct {
	db            *dbnary.DB
	languages     map[string]utils.TTLFile
	entryTemplate *template.Template
}

// ServeHTTP server up a dictionary entry
func (d *Dictionary) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) == 3 {
		lang, word := path[1], path[2]
		if _, ok := d.languages[lang]; !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		entry, err := d.db.LookupWordForLanguage(word, lang)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = d.entryTemplate.Execute(w, entry)
		if err != nil {
			return
		}
	}
}

// Server start server mode
func Server(db *dbnary.DB) {
	languages := make(map[string]utils.TTLFile)
	for _, language := range utils.TTLFiles {
		languages[language.Key] = language
	}
	supports := func(iso string) bool {
		_, ok := languages[iso]
		return ok
	}
	language := func(iso string) string {
		return utils.Languages[iso].Name
	}
	clean := func(word string) string {
		return strings.Replace(word, "_", " ", -1)
	}
	entryTemplate, err := template.New("entry").
		Funcs(map[string]interface{}{
			"supports": supports,
			"language": language,
			"clean":    clean,
		}).
		Parse(EntryTemplate)
	if err != nil {
		panic(err)
	}
	dictionary := Dictionary{
		db:            db,
		languages:     languages,
		entryTemplate: entryTemplate,
	}
	server := http.Server{
		Addr:    ":8080",
		Handler: &dictionary,
	}
	err = server.ListenAndServe()
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
	server = flag.Bool("server", false, "start up in server mode")
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

	if *server {
		Server(db)
		return
	}

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
