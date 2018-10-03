// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dbnary

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU(3)
	for i := 0; i < 12; i++ {
		_, ok := lru.Get(fmt.Sprintf("%d", i))
		if ok {
			t.Fatal("node should be not be found")
		}
		node := lru.Flush()
		if i < 7 && node != nil {
			t.Fatal("no nodes should be flushed", i)
		} else if i == 7 {
			if node == nil {
				t.Fatal("nodes should be flushed", i)
			} else {
				state, j := []string{"3", "2", "1", "0"}, 0
				for node != nil {
					if state[j] != node.Key {
						t.Fatal("state doesn't match")
					}
					j++
					node = node.B
				}
			}
		} else if i > 7 && i < 11 && node != nil {
			t.Fatal("no nodes should be flushed", i)
		} else if i == 11 {
			if node == nil {
				t.Fatal("nodes should be flushed", i)
			} else {
				state, j := []string{"7", "6", "5", "4"}, 0
				for node != nil {
					if state[j] != node.Key {
						t.Fatal("state doesn't match")
					}
					j++
					node = node.B
				}
			}
		}
	}

	check := func(key string, state []string) {
		if _, ok := lru.Get(key); !ok {
			t.Fatalf("key %s should be found", key)
		}
		node, i := lru.Head, 0
		t.Log("forward", state)
		for node != nil {
			t.Log(node.Key)
			if state[i] != node.Key {
				t.Fatal("invalid key", state[i], node.Key, key)
			}
			i++
			node = node.B
		}

		node, i = lru.Tail, len(state)
		t.Log("backward", state)
		for node != nil {
			i--
			t.Log(node.Key)
			if state[i] != node.Key {
				t.Fatal("invalid key", state[i], node.Key, key)
			}
			node = node.F
		}
	}
	check("8", []string{"8", "11", "10", "9"})
	check("10", []string{"10", "8", "11", "9"})
	check("8", []string{"8", "10", "11", "9"})
	check("8", []string{"8", "10", "11", "9"})

	lru = NewLRU(1)
	if _, ok := lru.Get("0"); ok {
		t.Fatal("should not find key 0")
	}
	if lru.Flush() != nil {
		t.Fatal("there shouldn't be a flush")
	}
	if _, ok := lru.Get("1"); ok {
		t.Fatal("should not find key 1")
	}
	if lru.Flush() == nil {
		t.Fatal("there should be a flush")
	}
	check("1", []string{"1"})
	check("1", []string{"1"})
	check("1", []string{"1"})
}
