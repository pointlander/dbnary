// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dbnary

import (
	"testing"
)

func TestSense(t *testing.T) {
	definitions := make(Definitions, 0, 8)
	definitions.Add("test1", "1")
	definitions.Add("test2", "1.1")
	definitions.Add("test3", "1a")
	definitions.Add("test4", "1,1")
	definitions.Add("test5", "1.1.1")
	definitions.Add("test6", "1.1.a")
	definitions.Add("test7", "10.13")
	definitions.Sort()
	correct := []string{"test1", "test2", "test4", "test5", "test6", "test3", "test7"}
	for i, definition := range definitions {
		if definition.Definition != correct[i] {
			t.Fatal("invalid sorting")
		}
	}
}
