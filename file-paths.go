// https://gobyexample.com/file-paths
//
// Parse & construct file paths.
// OS agnostic (handles Linux & Windows styles).

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// Use Join to build portable paths.
	// It takes any number of args, and constructs a hierarchical path.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Always use Join instead of concatenating.
	// It's portable, and normalises paths.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base can be used to get the directory path & filename.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Check whether a path is absolute.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// Handle extensions
	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fn2 := "goat.drawio.png"
	ext2 := filepath.Ext(fn2)
	fmt.Println(ext2)

	// Get the part without the extension
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a target.
	// It returns an error if that isn't possible.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
