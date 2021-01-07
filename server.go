// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dbnary

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/pointlander/dbnary/utils"
)

// ResultsTemplate is the template for search results
const ResultsTemplate = `<html>
 <head>
  <title>Search results for {{clean .Word}}</title>
  </head>
  <body>
{{$wordLanguage := .Language}}
		<ul>
{{range .Results}}
			<li><a href="/word/{{clean $wordLanguage}}/{{clean .Key}}">{{clean .Key}}</a></li>
{{end}}
		</ul>
  </body>
 </html>
`

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
	(<a href="/api/{{$wordLanguage}}/{{.Word}}">json</a>)<br/>
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
	db              *DB
	languages       map[string]utils.TTLFile
	entryTemplate   *template.Template
	resultsTemplate *template.Template
}

// Search searches for a word
func (d *Dictionary) Search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lang, word := ps.ByName("language"), ps.ByName("word")
	if _, ok := d.languages[lang]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	type Results struct {
		Language string
		Word     string
		Results  []Result
	}
	result, err := d.db.SearchWordForLanguage(word, lang)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	results := Results{
		Language: lang,
		Word:     word,
		Results:  result,
	}
	err = d.resultsTemplate.Execute(w, results)
	if err != nil {
		return
	}
}

// Word outputs the page for a word
func (d *Dictionary) Word(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lang, word := ps.ByName("language"), ps.ByName("word")
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

// API outputs the api representation of a word
func (d *Dictionary) API(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lang, word := ps.ByName("language"), ps.ByName("word")
	if _, ok := d.languages[lang]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	entry, err := d.db.LookupWordForLanguage(word, lang)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(entry)
	if err != nil {
		return
	}
}

// Server start server mode
func Server(db *DB, router *httprouter.Router) {
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

	resultsTemplate, err := template.New("entry").
		Funcs(map[string]interface{}{
			"supports": supports,
			"language": language,
			"clean":    clean,
		}).
		Parse(ResultsTemplate)
	if err != nil {
		panic(err)
	}

	dictionary := Dictionary{
		db:              db,
		languages:       languages,
		entryTemplate:   entryTemplate,
		resultsTemplate: resultsTemplate,
	}
	router.GET("/word-search/:language/:word", dictionary.Search)
	router.GET("/word/:language/:word", dictionary.Word)
	router.GET("/api/:language/:word", dictionary.API)
}
