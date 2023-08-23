// https://gobyexample.com/exit
//
// Immediately exit with a given status.

package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("!!!!")

	os.Exit(3)
}
