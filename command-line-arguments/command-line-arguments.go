// https://gobyexample.com/command-line-arguments
//
// Command line arguments!
// e.g. go run ./command-line-arguments/ bread cheese sausage herrings

package main

import (
	"fmt"
	"os"
)

func main() {

	// The raw arguments (including argv[0], the program path).
	argsWithProg := os.Args
	// The "actual" arguments.
	argsNoProg := os.Args[1:]

	// Get a specific argument (duh)
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsNoProg)
	fmt.Println(arg)
}
