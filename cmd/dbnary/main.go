// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/pointlander/dbnary"
)

var (
	print  = flag.String("print", "", "entry to print")
	depth  = flag.Int("depth", 0, "the depth to print entries")
	lookup = flag.String("lookup", "", "lookup a word")
)

func main() {
	flag.Parse()

	db := dbnary.OpenDB("dbnary.db", true)
	defer db.Close()

	if *print != "" {
		db.PrintEntry(*print, "", *depth, 0)
		return
	}

	if *lookup != "" {
		word, err := db.LookupWord(*lookup)
		if err != nil {
			panic(err)
		}
		data, err := json.MarshalIndent(word, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}
