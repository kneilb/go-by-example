// https://gobyexample.com/sorting
//
// There's a package for it. Yay.

package main

import (
	"fmt"
	"sort"
)

func main() {

	// Sort methods are specific to the builtin type.
	// This is an example for strings.
	// Sorting occurs in-place, so the slice gets changed & we don't return a new one.
	strs := []string{"cheese", "badger", "badge", "rancid", "Wolf"}
	fmt.Println("Strings:", strs)
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Sort ints.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// Check if a slice is already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted?!:", s)
}
