// https://gobyexample.com/command-line-flags
//
// Specify options for a program (e.g. wc -l)
//
// Go provides a flag package, which supports basic command-line parsing.

package main

import (
	"flag"
	"fmt"
)

func main() {

	// Go supports string, integer & boolean options.
	// This declares a string flag "word", with a default "foo" and some help text.
	// This function returns a pointer (not a value).
	wordPtr := flag.String("word", "foo", "a string")

	// This declares num and fork flags, using a similar approach.
	numPtr := flag.Int("num", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// You can also declare an option that uses an existing var.
	// NB. We have to pass in a pointer.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("num: ", *numPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
