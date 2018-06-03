// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"

	"github.com/pointlander/dbnary"
)

var printFlag = flag.String("print", "", "entry to print")

func main() {
	flag.Parse()

	db := dbnary.OpenDB("dbnary.db", true)
	defer db.Close()

	var display func(key, spaces string, depth int)
	display = func(key, spaces string, depth int) {
		fmt.Println(key)
		entry, err := db.GetEntry(key)
		if err != nil {
			panic(err)
		}
		translation, err := db.GetTranslation(key)
		if err != nil {
			panic(err)
		}

		printTerm := func(term *dbnary.Term) string {
			link := ""
			switch term.Type {
			case dbnary.Term_Blank:
				link = term.Key
				fmt.Print(term.Key)
			case dbnary.Term_IRI:
				prefix := dbnary.Prefixes[term.Prefix]
				fmt.Print(prefix.Name)
				fmt.Print(":")
				if prefix.Name == "eng" {
					link = term.Key
					fmt.Print(term.Key)
				} else {
					fmt.Print(prefix.Suffixes[term.Suffix])
				}
			case dbnary.Term_Literal:
				fmt.Print(term.Literal)
			}

			return link
		}
		if entry != nil {
			for _, trpl := range entry.Triples {
				fmt.Print(spaces)
				printTerm(trpl.Predicate)
				fmt.Print(" ")
				link := printTerm(trpl.Object)
				fmt.Print("\n")
				if link != "" && depth < 0 {
					display(link, spaces+" ", depth+1)
				}
			}
		}
		if translation != nil {
			for _, key := range translation.Keys {
				fmt.Print(spaces, " translation ")
				fmt.Println(key)
			}
		}
	}
	display(*printFlag, "", 0)
}
