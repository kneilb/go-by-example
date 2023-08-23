// https://gobyexample.com/defer
//
// Defer ensures that a function call is performed later.
// It's usually used to clean up.
// Kinda like finally / ensure in other languages.

package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("create")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("write")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("close")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
