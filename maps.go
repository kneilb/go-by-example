// https://gobyexample.com/maps

package main

import "fmt"

func main() {

	// Create an empty map using the "make" builtin
	m := make(map[string]int)
	fmt.Println("empty:", m)

	// Add k,v pairs using [] =, as per C++ etc
	m["k1"] = 7
	m["neil"] = 1979
	fmt.Println("with 2 keys:", m)

	// Read values using = [], as per C++ etc
	v1 := m["neil"]
	fmt.Println("read:", v1)

	// Non-existent values return the zero-value of the type (uh oh)
	v2 := m["fishface"]
	fmt.Println("read nonexistent:", v2)

	// len() returns number of k,v pairs
	fmt.Println("len:", len(m))

	// Delete a key
	delete(m, "k1")
	fmt.Println("after delete:", m)

	// Deleting non-existent is not an error
	// No return value to indicate whether something is deleted or not
	delete(m, "k1")

	// The optional 2nd return value indicates whether the key was present when reading
	_, present := m["k2"]
	fmt.Println("present:", present)

	// Can initialise a map in a single line
	n := map[string]int{"fish": 1, "whale": 12, "hake": 9}
	fmt.Println("init:", n)
}
