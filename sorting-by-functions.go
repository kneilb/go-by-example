// https://gobyexample.com/sorting-by-functions
//
// Sort a collection by something other than its "natural order".
// In other words, custom sorting.

package main

import (
	"fmt"
	"sort"
)

// We need a type that implements sort.Interface
// This type is just an alias for the builtin []string.
type byLength []string

// Len() and Swap() will be the same as for []string.
// The interesting bit is Less(), which has the custom sorting logic.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Do the custom sort by converting fruits into a byLength, then Sort()ing that.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
