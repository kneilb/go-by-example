// https://gobyexample.com/string-formatting
//

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Includes struct field names
	fmt.Printf("struct2: %+v\n", p)

	// Go syntax representation, i.e. source code snippet!
	fmt.Printf("struct3: %#v\n", p)
	// The type of the value
	fmt.Printf("type: %T\n", p)

	// boolean
	fmt.Printf("bool: %t\n", true)

	// integer in decimal (base 10)
	fmt.Printf("int: %d\n", 14)

	// integer in binary (base 2)
	fmt.Printf("bin: %b\n", 14)

	// integer as character (ASCII)
	fmt.Printf("char: %c\n", 33)

	// integer in hexadecimal (base16)
	fmt.Printf("hex: %x\n", 456)

	// float as basic decimal
	fmt.Printf("float1: %f\n", 78.9)

	// float in scientific notation
	fmt.Printf("float2: %e\n", 123400000.0)
	// float in slightly different SN
	fmt.Printf("float3: %E\n", 123400000.0)

	// basic string print
	fmt.Printf("str1: %s\n", "\"string\"")

	// print quoted string, as in Go source
	fmt.Printf("str2: %q\n", "\"string\"")

	// print string as hex digits, 2 chars out per byte of input
	fmt.Printf("str3: %x\n", "hex this")

	// pointer
	fmt.Printf("pointer: %p\n", &p)

	// specify width - right justify, pad with spaces by default
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// specify width of float - width.precision syntax
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// specify width - left align
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// specify width of string - right justify by default
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	// left justify using -, as normal
	fmt.Printf("width4: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
