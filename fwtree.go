// Copyright 2014 Mahmud Ridwan. All rights reserved.

// Package fwtree provides an implementation of Fenwick tree (http://en.wikipedia.org/wiki/Fenwick_tree).
//
// A Fenwick tree is a data structure that provides efficient methods for calculation and manipulation of the prefix sums of a table of values.
package fwtree

// Fwtree represents a Fenwick tree. The zero value for Fwtree is an tree of size 0, not very useful.
type Fwtree struct {
	t []int64
	n int
}

func New(n int) *Fwtree {
	return &Fwtree{make([]int64, n+2), n}
}

// Len returns the size of the Fenwick tree.
func (f Fwtree) Len() int {
	return f.n
}

// Add adds v to index i.
//
// Add panics when i < 0 or i >= f.Len().
func (f Fwtree) Add(i int, v int64) {
	if i < 0 || i >= len(f.t) {
		panic("fwtree: index out of range")
	}

	i++
	for i < len(f.t) {
		f.t[i] += v
		i += i & -i
	}
}

// Sum returns the sum of all values starting from index 0 up to r (exclusive).
//
// Sum panics when r < 0 or r > f.Len().
func (f Fwtree) Sum(r int) int64 {
	if r < 0 || r > len(f.t) {
		panic("fwtree: index out of range")
	}

	var z int64
	for r > 0 {
		z += f.t[r]
		r -= r & -r
		if r == 0 {
			break
		}
	}
	return z
}

// At returns the value at index i.
//
// At panics when i < 0 or i >= f.Len().
func (f Fwtree) At(i int) int64 {
	return f.Sum(i+1) - f.Sum(i)
}
