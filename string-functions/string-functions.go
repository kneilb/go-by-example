// https://gobyexample.com/string-functions
//
// Examples of standard library string package functions

package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1)) // Replace all
	p("Replace:   ", s.Replace("foo", "o", "0", 1))  // Replace 1
	p("Split:     ", s.Split("a-b-c-d-e-f-", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
