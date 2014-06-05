// Copyright 2014 Mahmud Ridwan. All rights reserved.

package fwtree_test

import (
	"fmt"

	"github.com/hjr265/fwtree.go"
)

func ExampleFwtree() {
	// Create a Fenwick tree of size 32.
	f := fwtree.New(32)

	f.Add(9, 14)
	fmt.Println(f.Sum(5))
	fmt.Println(f.Sum(9))
	fmt.Println(f.Sum(32))

	f.Add(26, 0)
	fmt.Println(f.Sum(5))
	fmt.Println(f.Sum(12))
	fmt.Println(f.Sum(32))

	f.Add(1, 5)
	fmt.Println(f.Sum(5))

	f.Add(31, 6)
	fmt.Println(f.Sum(12))
	fmt.Println(f.Sum(32))

	// Output:
	// 0
	// 0
	// 14
	// 0
	// 14
	// 14
	// 5
	// 19
	// 25

}
