// https://gobyexample.com/rate-limiting
//
// Rate limiting allows us to control resource utilisation & maintain QoS.
// Go supports rate limiting with goroutines, channels & tickers.

package main

import (
	"fmt"
	"time"
)

// Basic rate limiting.
// Suppose we want to limit our handling of incoming requests.
func basic() {
	fmt.Println("basic")

	// Serve these requests off a channel called "requests"
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// The limiter channel will get a value every 200 milliseconds.
	// We can use this to regulate things.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each request
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}

// Bursty rate limiting.
// Allow short bursts of requests, but still preserve the overall limit.
// We can do this by buffering the limiter channel!
func bursty() {
	fmt.Println("bursty")

	// This burstyLimiter channel allows bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel so we can start with a burst of 3
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Add another entry to the burstyLimiter channel every 200ms
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests.
	// The first 3 will be "bursted" through immediately, the last 2 will be delayed by 200ms.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func main() {
	// This batch of requests is handled every 200ms
	basic()

	// In this batch, the first 3 are handled immediately. The remaining 2 each have a 200ms delay.
	bursty()
}
