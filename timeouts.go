// https://gobyexample.com/timeouts
//
// Timeouts are important, m'kay!
// They're needed when connecting to external resources that might take ages,
// or generally if we need to bind our execution time.
// Timeouts in go are lovely, with channels and select.

package main

import (
	"fmt"
	"time"
)

func main() {

	// An external call returns its result on c1 after 2 seconds.
	// The channel is buffered, so the send is non-blocking.
	// This is a common pattern to prevent goroutine leaks, in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// This select implements a timeout
	select {
	// Wait for the result (await, normal)
	case res := <-c1:
		fmt.Println(res)
	// Awaits a value that gets sent after a 1 second timeout
	// This will get taken if the operation takes more than the allowed 1 second (which it will...!)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If we allow a longer timeout, then the receive succeeds, and we print the result
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
