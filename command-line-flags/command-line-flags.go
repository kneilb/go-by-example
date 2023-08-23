// https://gobyexample.com/command-line-flags
//
// Specify options for a program (e.g. wc -l)
//
// Go provides a flag package, which supports basic command-line parsing.
//
// Build it using "go build command-line-flags.go"
//
// Provide all values
//  ./command-line-flags -word=opt -numb=7 -fork -svar=flag
//
// Omitted flags take their defaults
//  ./command-line-flags -word=opt
//
// Trailing positional arguments can be provided after flags
//  ./command-line-flags -word=opt a1 a2 a3
//
// The flag package requires all flags to appear before positional arguments!
// Otherwise the flags are interpreted as position arguments...
//  ./command-line-flags -word=opt a1 a2 a3 -numb=7
//
// Use --help or -h to get automatically generated help text.
//  ./command-line-flags -h
//
// If you provide a flag that wasn't specified, you get an error message & the help text.
//  ./command-line-flags --fishface

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

	// Once it's all declared, run the parser.
	flag.Parse()

	// Print out the values we've parsed.
	fmt.Println("word:", *wordPtr)
	fmt.Println("num: ", *numPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
