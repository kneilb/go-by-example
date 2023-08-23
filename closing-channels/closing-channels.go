// https://gobyexample.com/closing-channels
//
// Closing a channel indicates that no more values will be sent on it.
// This can communicate completion to the receivers.

package main

import "fmt"

// We use a jobs channel to communicate work to be done from the main() goroutine to a separate worker goroutine.
// When there are no more jobs, we close the jobs channel.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// This is the worker goroutine.
	// It repeatedly receives from jobs, using the 2-value form of receive.
	// In this form, the "more" value will be false if the channel has been closed, and all values have been received.
	// We use this to notify on done when we've worked on all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Send 3 jobs to the worker over the jobs channel, then close it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Await the worker using the synch approach we used before
	<-done
}
