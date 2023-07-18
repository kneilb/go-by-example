// https://gobyexample.com/recursion

package main

import "fmt"

// Define a recursive function
// NAB: Make sure you have a robust termination condition
// NAB: This example doesn't - call fact(-1) for a demo ;)
func fact(n int) int {
	// NAB: fix with <= 0 or unsigned
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	// Don't do this, the function isn't very safe
	// fmt.Println(fact(-1))
	fmt.Println(fact(7))

	// Closures can be recursive, too
	// Although you have to forward-declare them!
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
