// https://gobyexample.com/tickers
//
// Tickers allow you to do something repeatedly at regular intervals

package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickers are similar to Timers.
	// They have a channel that is sent values.
	// Here we use select on the channel to await values that arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped, like Timers.
	// Once stopped we won't receive any more values on its channel.
	// We stop ours after 1600ms, so the ticker should tick 3 times.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
