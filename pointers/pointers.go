// https://gobyexample.com/pointers

package main

import "fmt"

// C-style value syntax
func zeroval(ival int) {
	ival = 0
}

// C-style pointer syntax
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// Just like C; i passed by value can't be modified
	zeroval(i)
	fmt.Println("zeroval(i):", i)

	// Again, just like C; i passed as a pointer can be modified
	zeroptr(&i)
	fmt.Println("zeroptr(i):", i)

	// Addresses can be printed, too
	fmt.Println("pointer:", &i)
}
