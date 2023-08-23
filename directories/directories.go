// https://gobyexample.com/directories
//

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

	// Create a new subdirectory in the current working dir.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Good practice to defer their removal.
	// RemoveAll will delete a whole tree (like rm -rf).
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Create a hierarchy of directories - like mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of os.DirEntry objects.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir lets us change working directory.
	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to where we started.
	err = os.Chdir("../../..")
	check(err)

	// We can also visit a directory recursively, including all its sub-directories.
	// Walk accepts a callback function to handle the visited files/directories.
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// This function is called for every file or directory found recursively by filepath.Walk.
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
