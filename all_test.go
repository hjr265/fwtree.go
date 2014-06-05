// Copyright 2014 Mahmud Ridwan. All rights reserved.

package fwtree_test

import (
	"testing"

	"github.com/hjr265/fwtree.go"
)

type fetree []int64

func (f fetree) Add(i int, v int64) {
	f[i] += v
}

func (f fetree) Sum(r int) int64 {
	var z int64
	for i := 0; i < r; i++ {
		z += f[i]
	}
	return z
}

func (f fetree) At(i int) int64 {
	return f[i]
}

func TestFwtree(t *testing.T) {
	ts := []struct {
		i int
		v int64
	}{
		{2, 19},
		{7, 3},
		{18, 9},
		{9, 14},
		{26, 0},
		{1, 5},
		{28, 1},
		{31, 6},
		{19, 2},
		{0, 19},
		{3, 7},
		{9, 18},
		{14, 9},
		{1, 28},
		{6, 0},
		{6, 31},
	}
	fe := make(fetree, 32)
	fw := fwtree.New(32)

	for _, tc := range ts {
		fe.Add(tc.i, tc.v)
		fw.Add(tc.i, tc.v)

		for r := 0; r <= 32; r++ {
			e := fe.Sum(r)
			s := fw.Sum(r)
			if e != s {
				t.Fatalf("Sum(%d) = %d, want %d", r, s, e)
			}

			if r == 32 {
				continue
			}
			e = fe.At(r)
			s = fw.At(r)
			if e != s {
				t.Fatalf("At(%d) = %d, want %d", r, s, e)
			}
		}
	}
}
