// https://gobyexample.com/variables

package main

import "fmt"

func main() {
	var a = "Initial"
	fmt.Println("a string", a)
	a += " 1"
	fmt.Println("a modified string", a)

	// structured bindings
	var b, c int = 1, 2
	fmt.Println("structured bindings b, c", b, c)

	var d = true
	fmt.Println("true boolean value d", d)

	var e int
	fmt.Println("uninitialised int e", e)

	// shorthand for
	// var f string = "apple"
	f := "shorthand initialisation, apple"
	fmt.Println("f", f)
}
