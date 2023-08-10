// https://gobyexample.com/regular-expressions
//
// Go has built-in support for regular expressions.
// Full docs are here: https://pkg.go.dev/regexp

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Previously we used the string pattern directly.
	// For other regexp tasks, you'll need to Compile an optimised Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// These structs support many methods. This is an equivalent MatchString test.
	fmt.Println(r.MatchString("peach"))

	// This fails because the regexp needs 1 or more letters between the p & the ch.
	fmt.Println(r.MatchString("pch"))

	// This returns the matched string.
	fmt.Println(r.FindString("peach punch"))

	// This returns the starting and ending indices, rather than the actual string.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// The Submatch variants include information about both the whole-pattern matches and the submatches.
	// Here, this returns information for both p([a-z]+)ch) and ([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similarly, this returns indices for the submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The All variants apply to all matches in the input, not just the first (like sed -g).
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// There are All variants of the other functions, too
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Providing a non-negative integer as the 2nd argument limits the number of matches.
	fmt.Println("all:", r.FindAllString("peach punch pinch", 2))

	// There are also versions that take []byte arguments.
	fmt.Println(r.Match([]byte("peach")))

	// When creating global variables with regexps, you can use MustCompile.
	// This will panic instead of returning an error.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// Replace subsets of strings with other values.
	// We're replacing the "bit in brackets" (AKA marked subexpression, block, capturing group).
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func variant allows us to transform matched text using a function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
