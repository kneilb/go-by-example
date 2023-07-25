// https://gobyexample.com/select
//
// select lets you wait on multiple channel operations
// combining goroutines, channels & select is good, m'kay
// like old skool POSIX synchronisation primitives, built into the language

package main

import (
	"fmt"
	"time"
)

func main() {

	// 2 channels
	c1 := make(chan string)
	c2 := make(chan string)

	// 2 goroutines, which send data on the channels after a delay
	// This could be blocking RPC operations, HTTP long polling, etc etc
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We use select to wait for EITHER of these to arrive, in whichever order
	// Execution time is ~2 seconds, because the two sleeps occur concurrently
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
