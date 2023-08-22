// https://gobyexample.com/temporary-files-and-directories
//
// These are often needed, and jolly good - hurray!

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Easy version!
	// Creates a file, and opens it for R+W.
	// "" as a first arg means to create it in the default location.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Display the name of the temporary file.
	// On Unix-based OS, this is likely to be in /tmp.
	// The file name starts with the prefix that's provided as the 2nd argument.
	// The rest is automatic, to ensure concurrent accesses are safe.
	fmt.Println("Temp file name:", f.Name())

	// Clean up explicitly.
	// The OS will do this at some point, but it's better to do it ourselves.
	defer os.Remove(f.Name())

	// Write some data to the file.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// If we're expecting to create lots of temp files, we can create a temp dir instead.
	// This gives a directory name, rather than an open file.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Now we create temp file names by prefixing them with the temp directory name.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
