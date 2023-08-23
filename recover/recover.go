// https://gobyexample.com/recover
//
// We can recover from a panic, using the recover built-in.
// It stops the panic aborting the whole program, and continue instead.
// Kinda like catching an exception...
//
// This is useful if you have a web server, and a goroutine handling a single request panics.
// In this case we'd want to just kill that request and carry on with the others.
// I've tried to demo something like that below...!

package main

import (
	"fmt"
	"time"
)

// A very panicky function.
func mayPanic() {
	fmt.Println("During mayPanic()")
	panic("poblem!")
}

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			// The defer/recover has to be inside each goroutine.
			// Putting it in main doesn't help.
			defer func() {
				// The return value of recover is nil if no error occurred, or the error raised in any panic()
				if r := recover(); r != nil {
					fmt.Println("Recovered. Error:\n", r)
				}
			}()

			mayPanic()

			// We never get here, execution ends inside the defer closure.
			fmt.Println("After mayPanic()")
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("Main after go")
}
