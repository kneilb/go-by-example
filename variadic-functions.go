// https://gobyexample.com/variadic-functions

package main

import "fmt"

// This function takes an arbitrary number of int variables
func sum(nums ...int) {
	// Inside the function, the argument behaves like []int
	fmt.Println(nums)

	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum()
	sum(0)

	sum(1)

	sum(1, 2, 3)

	nums := []int{1, 2, 7, 19, 27, 44}
	sum(nums...)
}
