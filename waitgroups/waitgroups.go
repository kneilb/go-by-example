// https://gobyexample.com/waitgroups
//
// WaitGroups can be used to wait for multiple goroutines to finish.

package main

import (
	"fmt"
	"sync"
	"time"
)

// The worker function we run in all the goroutines.
// Sleep to simulate an expensive task.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// WaitGroup used to wait for all the goroutines to finish.
	// NB. If a WaitGroup is passed to a function, it should be done BY POINTER!
	var wg sync.WaitGroup

	// Launch several goroutines, and increment the WaitGroup counter by one for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// avoid reuse of the same i value in each goroutine closure.
		// this is ugly; see here for an explanation: https://go.dev/doc/faq#closures_and_goroutines
		// essentially it's captured by pointer :S
		i := i

		// Wrap the worker call in a closure, which tells the WaitGroup that the worker is done.
		// This means that the worker doesn't have to know about the concurrency primitives we're using.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Block until the WaitGroup counter goes back to 0 (as the workers notify that they are done)
	// There is no straightforward way to pass up errors from workers.
	// For a more advanced approach, try the errgroup package (https://pkg.go.dev/golang.org/x/sync/errgroup)
	wg.Wait()
}
