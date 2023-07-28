// https://gobyexample.com/panic
//
// Used to terminate the program if something unexpected goes wrong.

package main

import "os"

func main() {

	// Generally used to deal with unexpected errors.
	// This means that we'll just panic on run.
	// Ideally we'll use error-indicating return values instead.
	// panic("sum ting wong")

	// Example of panicking if we can't create a file.
	// This code won't get executed though; we already panicked!
	_, err := os.Create("/tmp/fish/file")
	if err != nil {
		panic(err)
	}
}
