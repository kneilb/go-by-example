// https://gobyexample.com/environment-variables
//
// Environment variables are a way to get config info into Unix programs.
// This example explores setting, getting & listing them.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Set a key/value pair
	os.Setenv("FOO", "1")

	// Get one
	fmt.Println(os.Getenv("FOO"))

	// Unset variables return an empty string.
	// If we set BAR before running the program, we'll get that value.
	fmt.Println(os.Getenv("BAR"))

	// os.Environ() lists all the key/value pairs in the environment.
	// It returns a slice of strings in the form KEY=value.
	// Use strings.SplitN to separate them.
	// This prints out all the keys - what's printed here is machine dependent.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
