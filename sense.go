// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dbnary

import (
	"sort"
	"unicode"
)

var _ sort.Interface = &Definitions{}

// SenseState is a sense parsing state
type SenseState int

const (
	// SenseStateUnknown is an uknown state
	SenseStateUnknown SenseState = iota
	// SenseStateDigit is a digit state
	SenseStateDigit
	// SenseStateLetter is a letter state
	SenseStateLetter
)

// Definition is a definition with a sense ordinal
type Definition struct {
	Definition string
	Sense      []string
	Ordinal    string
}

// Definitions is a list of definition
type Definitions []Definition

// Add adds a defintion
func (d *Definitions) Add(definition, sense string) {
	newDefinition := Definition{
		Definition: definition,
	}
	state, token := SenseStateUnknown, ""
	for _, s := range sense {
		isDigit, isLetter := unicode.IsDigit(s), unicode.IsLetter(s)
		switch state {
		case SenseStateUnknown:
			switch true {
			case isDigit:
				state = SenseStateDigit
				token += string(s)
			case isLetter:
				state = SenseStateLetter
				token += string(s)
			}
		case 1:
			switch true {
			case isDigit:
				token += string(s)
			case isLetter:
				newDefinition.Sense = append(newDefinition.Sense, token)
				state = SenseStateLetter
				token = string(s)
			default:
				newDefinition.Sense = append(newDefinition.Sense, token)
				state = SenseStateUnknown
				token = ""
			}
		case 2:
			switch true {
			case isDigit:
				newDefinition.Sense = append(newDefinition.Sense, token)
				state = SenseStateDigit
				token = string(s)
			case isLetter:
				token += string(s)
			default:
				newDefinition.Sense = append(newDefinition.Sense, token)
				state = SenseStateUnknown
				token = ""
			}
		}
	}
	if token != "" {
		newDefinition.Sense = append(newDefinition.Sense, token)
	}
	*d = append(*d, newDefinition)
}

// Normalize prepases the definitions for sorting
func (d *Definitions) Normalize() {
	maxLength, max := 0, 0
	for _, definition := range *d {
		if length := len(definition.Sense); length > maxLength {
			maxLength = length
		}
		for _, sense := range definition.Sense {
			if length := len([]rune(sense)); length > max {
				max = length
			}
		}
	}
	filler := ""
	for i := 0; i < max; i++ {
		filler += " "
	}
	for i, definition := range *d {
		for _, sense := range definition.Sense {
			fill := max - len(sense)
			for j := 0; j < fill; j++ {
				(*d)[i].Ordinal += " "
			}
			(*d)[i].Ordinal += sense
		}
	}
}

// Len is the length of the definitions
func (d *Definitions) Len() int {
	return len(*d)
}

// Swap swaps to definitions
func (d *Definitions) Swap(i, j int) {
	(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
}

// Less determines if one ordinal is less than the other
func (d *Definitions) Less(i, j int) bool {
	return (*d)[i].Ordinal < (*d)[j].Ordinal
}

// Sort sorts the definitions by sense
func (d *Definitions) Sort() {
	d.Normalize()
	sort.Sort(d)
}
