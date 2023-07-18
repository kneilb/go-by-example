// https://gobyexample.com/range

package main

import "fmt"

func main() {
	// Create a slice, and sum it with a range
	numbers := []int{2, 3, 4}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Get the index out of the range, as well as the value
	for i, num := range numbers {
		if num == 3 {
			fmt.Println("Found number", num, "at index", i)
		}
	}

	// range on a map iterates over key / value pairs
	m := map[string]int{"fish": 1, "cheese": 2, "ninja": 3}
	for k, v := range m {
		fmt.Println("map entry:", k, "=", v)
	}

	// range can also iterate over just the map keys
	for k := range m {
		fmt.Println("map key:", k)
	}

	// Iterates over Unicode code points
	// First value is starting byte index of the rune
	// Second is the rune itself
	for i, c := range "go" {
		fmt.Println("string:", i, c)
	}

}
