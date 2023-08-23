// https://gobyexample.com/stateful-goroutines
//
// This is a much better way of achieving the same thing as mutexes.go
// It aligns with Go's ideas of sharing memory by communicating,
// and having each piece of data owned by exactly 1 goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// State is owned by a single gouroutine.
// This guarantees that the data isn't corrupted by concurrent access.
// To read and write the state, other goroutines send messages to the owner, and receive replies.
// The readOp and writeOp structs encapsulate the requests, and a way for the owner to respond.
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Counters, to see how many reads & writes we perform
	var readOps uint64
	var writeOps uint64

	// These channels are used by goroutines to issue read & write requests
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// This is the owner goroutine.
	// It owns state, which is map that it holds privately.
	// It repeatedly selects on the read & write channels, responding to requests as they arrive.
	// It performs the selected operation, then sends a response on the resp channel.
	// The response contains the read value for readOps.
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Start 100 goroutines that request reads.
	// Each read constructs a readOp, including the unique resp channel.
	// It sends the readOp over the reads channel, and receives the response via resp.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 goroutines requesting writes.
	// The approach is similar, but we send the value to read
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Run the goroutines for a second.
	time.Sleep(time.Second)

	// Capture and report the number of operations.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// When run, we get approximately 80k ops in 1 second.
	// This approach seems like overkill for this case - it's way more complicated than the mutex-based approach.
	// It may make more sense if:
	// - other channels are involved
	// - multiple mutexes would be needed, which is error-prone
	//
	// Use whichever feels more natural, man.
}
