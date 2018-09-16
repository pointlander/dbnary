// Copyright 2018 The dbnary Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dbnary

// Node is an entry in the LRU cache
type Node struct {
	F, B  *Node
	Entry Entry
	Key   string
	Seen  bool
}

// LRU is a least recently used cache
type LRU struct {
	Size       int
	Head, Tail *Node
	Nodes      map[string]*Node
}

// NewLRU creates a new LRU cache
func NewLRU(size uint) LRU {
	if size == 0 {
		panic("size should not be 0")
	}
	return LRU{Size: 1 << size}
}

// Flush flush the oldest entries in the cache
func (l *LRU) Flush() *Node {
	size := l.Size
	if len(l.Nodes) < size {
		return nil
	}

	node := l.Tail
	delete(l.Nodes, node.Key)
	size >>= 1
	for i := 1; i < size; i++ {
		node = node.F
		delete(l.Nodes, node.Key)
	}
	node.F.B, l.Tail, node.F = nil, node.F, nil

	return node
}

// Get gets an entry and sets it as the most recent
func (l *LRU) Get(key string) (*Node, bool) {
	length := len(l.Nodes)
	if length > 0 {
		if node := l.Nodes[key]; node != nil {
			if node.F != nil {
				if node.B != nil {
					node.B.F, node.F.B = node.F, node.B
				} else {
					node.F.B, l.Tail = nil, node.F
				}
				node.F, node.B, l.Head, l.Head.F = nil, l.Head, node, node
			}

			return node, true
		}
	}

	node := &Node{Key: key}
	node.B, l.Head = l.Head, node
	if length == 0 {
		l.Tail = node
		l.Nodes = make(map[string]*Node, l.Size)
	} else {
		node.B.F = node
	}
	l.Nodes[key] = node
	return node, false
}
