// https://gobyexample.com/if-else

package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is even")
	}

	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num has one digit")
	} else {
		fmt.Println("num has more than one digit")
	}

	// There's no ternary if...! ?:(
}
