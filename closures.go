// https://gobyexample.com/closures

package main

import "fmt"

// This function returns a function,
// which is defined in the body of intSeq
// The returned function "closes over" (captures) the variable i
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// The state is unique - this is a new closure
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}
