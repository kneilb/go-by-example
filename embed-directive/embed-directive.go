// https://gobyexample.com/embed-directive
//
// go:embed is a compiler directive that allows programs to include arbitrary files & folder in the binary.
// Read more here: https://pkg.go.dev/embed
//
// mkdir -p folder
// echo "hello go" > folder/single_file.txt
// echo "123" > folder/file1.hash
// echo "456" > folder/file2.hash
// go run embed-directive.go

package main

// Import the embed package.
// If you don't use any identifiers from this package (i.e. you only use the compiler directives),
// you can do a blank import with _ "embed"
import "embed"

// Embed directives accept paths relative to the directory containing the source file.
// This embeds the file contents into the string variable immediately following it.
//
//go:embed folder/single_file.txt
var fileString string

// This embeds the contents into a []byte.
//
//go:embed folder/single_file.txt
var fileByte []byte

// We can also embed multiple files, or folders with wildcards.
// This uses an embed.FS varible, which implements a virtual file system.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Print out the contents of single_file.txt
	print(fileString)
	print(string(fileByte))

	// Retrieve some files from the embedded folder.
	content, _ := folder.ReadFile("folder/file1.hash")
	print(string(content))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
