// https://gobyexample.com/writing-files
//
// Similar to reading!

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Dump a string (or just bytes...) to a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// To be more granular, open a file for writing.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// It's idiomatic to defer a close immediately after opening.
	defer f.Close()

	// Write byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Write strings
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Flush writes to stable storage.
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied.
	// Without this, the data isn't written (even on program exit!)
	w.Flush()
}
