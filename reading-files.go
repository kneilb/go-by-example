// https://gobyexample.com/reading-files
//
// Examples of reading files.
//
// echo "hello" > /tmp/dat
// echo "go" >> /tmp/dat
// go run reading-files.go

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Function for error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Shortcut for slurping a file's contents into memory.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))

	// FOr more control, open a file first to get an os.File value.
	f, err := os.Open("/tmp/dat")
	check(err)

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read, but remember how many were actually read!
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek to a known location, and Read from there.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// The io package provides some useful functions.
	// ReadAtLeast implements what we've done above, but more robustly!
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There's no Rewind, but Seek(0, 0) does the same thing.
	_, err = f.Seek(0, 0)
	check(err)

	// The bufio package provides a buffered reader.
	// It is more efficient for small reads, and implements more reading methods.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
