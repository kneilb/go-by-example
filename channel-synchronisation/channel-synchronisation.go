// https://gobyexample.com/channel-synchronization
//
// We can use a channel to synchronise between goroutines
// Here we're using a blocking receive to wait for a goroutine to finish
// To wait for multiple goroutines, it might be better to use a WaitGroup

package main

import (
	"fmt"
	"time"
)

// A worker function, which uses the passed "done" channel to indicate completion
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done!")

	done <- true
}

func main() {
	// make a channel
	done := make(chan bool, 1)

	// start a worker, giving it the channel to notify on
	go worker(done)

	// block until we receive notification from the worker
	// if we remove this, the program will exit before the worker even starts...!
	<-done
}
