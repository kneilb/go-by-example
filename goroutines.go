// https://gobyexample.com/goroutines

package main

import (
	"fmt"
	"time"
)

// Define an exciting function, f
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond)
	}
}

func main() {

	// Call f synchronously (i.e. in "normal" way)
	f("directly")

	// Run the function in a goroutine
	// It will execute concurrently with the calling function
	go f("goroutine")

	// Call an anonymous function in a goroutine
	go func(msg string) {
		for f := 0; f < 3; f++ {
			fmt.Println(msg, f)
			time.Sleep(time.Millisecond)
		}
	}("boing")

	// Our two functions are running asynchronously, in separate goroutines
	// Wait for them both to finish...!
	// This is a rubbish way of doing it - use a WaitGroup to do it properly! (see later exercise)
	time.Sleep(time.Second)
	fmt.Println("done")
}
