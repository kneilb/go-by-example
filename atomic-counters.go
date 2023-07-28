// https://gobyexample.com/atomic-counters
//
// We prefer to use communication over channels to manage state between multiple goroutines.
// There are a few other options though.
// This example shows how to use atomic counters...

package main

import (
	"fmt"
	"sync"
)

func main() {

	// This unsigned int represents our (positive) counter
	var ops uint64

	// This WaitGroup lets us wait for all the goroutines to finish
	var wg sync.WaitGroup

	// Start 50 goroutines that each increment the counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// Use AddUint64 to atomically increment the counter
		// We pass it the memory address of the counter, ops
		go func() {
			for c := 0; c < 1000; c++ {
				// atomic.AddUint64(&ops, 1)
				ops++
			}
			wg.Done()
		}()

		// Wait until all the goroutines are done
		wg.Wait()

		// We can access ops safely now; no other goroutines are running
		// We could use a function like atomic.LoadUint64 to read it safely even while others are running
		fmt.Println("ops:", ops)

		// We expect to get exactly 50k operations
		// If we'd used a non-atomic ops++ to increment the counter, we'd get a non-deterministic result
		// We'd also get data race failures when running with -race
	}
}
