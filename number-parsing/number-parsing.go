// https://gobyexample.com/number-parsing
//
// Parsing strings from numbers. Hurray!

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 64 tells PF how many bits of precision to use.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 64 means that the number must fit in 64 bits.
	// 0 means "work out the base for yourself"
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// Auto-parse hex formatted numbers with 0 base.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Unsigned version also available.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Shorthand for base-10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Return errors on bad input.
	_, e := strconv.Atoi("fish")
	fmt.Println(e)
}
