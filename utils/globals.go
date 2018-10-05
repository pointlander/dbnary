// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

// TTLFile is a ttl file
type TTLFile struct {
	Name, Key, Full string
}

var (
	// TTLFiles are the dbnary files
	TTLFiles = []TTLFile{
		{"bg_dbnary_ontolex.ttl.bz2", "bul", "bulgarian"},
		{"de_dbnary_ontolex.ttl.bz2", "deu", "german"},
		{"el_dbnary_ontolex.ttl.bz2", "ell", "greek"},
		{"en_dbnary_ontolex.ttl.bz2", "eng", "english"},
		{"es_dbnary_ontolex.ttl.bz2", "spa", "spanish"},
		{"fi_dbnary_ontolex.ttl.bz2", "fin", "finnish"},
		{"fr_dbnary_ontolex.ttl.bz2", "fra", "french"},
		{"id_dbnary_ontolex.ttl.bz2", "ind", "indonesian"},
		{"it_dbnary_ontolex.ttl.bz2", "ita", "italian"},
		{"ja_dbnary_ontolex.ttl.bz2", "jpn", "japanese"},
		{"la_dbnary_ontolex.ttl.bz2", "lat", "latin"},
		{"lt_dbnary_ontolex.ttl.bz2", "lit", "lithuanian"},
		{"mg_dbnary_ontolex.ttl.bz2", "mlg", "malagasy"},
		{"nl_dbnary_ontolex.ttl.bz2", "nld", "dutch"},
		{"no_dbnary_ontolex.ttl.bz2", "nor", "norwegian"},
		{"pl_dbnary_ontolex.ttl.bz2", "pol", "polish"},
		{"pt_dbnary_ontolex.ttl.bz2", "por", "portuguese"},
		{"ru_dbnary_ontolex.ttl.bz2", "rus", "russian"},
		{"sh_dbnary_ontolex.ttl.bz2", "hbs", "serbo-croatian"},
		{"sv_dbnary_ontolex.ttl.bz2", "swe", "swedish"},
		{"tr_dbnary_ontolex.ttl.bz2", "tur", "turkish"},
	}
)
