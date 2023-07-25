// https://gobyexample.com/worker-pools
//
// Implement a worker pool using goroutines & channels

package main

import (
	"fmt"
	"time"
)

// Worker function
// We run several concurrent instances in goroutines.
// They receive work on the jobs channel, and send results on the results channel.
// We sleep to simulate an "expensive" task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// To use our worker pool, we send them work and collect the results.
	// Make 2 channels for this.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start up 3 workers, initially blocked (there are no jobs yet).
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs and then close the channel, indicating that we have no more work.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect the results.
	// This also ensures that the workers goroutines have finished.
	// Alternatively we could use a WaitGroup... (see the next episode)
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	// When we run it, we get the 5 jobs being run by the various workers.
	// It takes about 2 seconds, despite doing 5 seconds worth of work.
	// This is because there are 3 workers running concurrently.
}
