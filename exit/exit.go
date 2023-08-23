// https://gobyexample.com/exit
//
// Immediately exit with a given status.

package main

import (
	"fmt"
	"os"
)

func main() {

	// defers will NOT be run when using os.Exit, so this will never happen.
	defer fmt.Println("!!!!")

	// Exit with status 3
	os.Exit(3)
}
