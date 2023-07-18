// https://gobyexample.com/strings-and-runes

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Literal "hello" in Thai
	const s = "สวัสดี"

	fmt.Println("string:", s)

	// This returns the size in bytes
	fmt.Println("len:", len(s))

	// This iterates over the bytes in the underlying array (= code points in s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// In Thai, some runes are diacritical marks, so this is 6 rather than 4!
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// This iterates over the runes in the utf8 string
	// (i.e. range handles strings in a "special" way)
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println()

	// This achieves the same as above, but using DecodeRuneInString explicitly
	// It's a lot wordier(!)
	fmt.Println("Using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
