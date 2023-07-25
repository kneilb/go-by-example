// https://gobyexample.com/timers
//
// Timers allow us to execute code at some point in the future.
//
// There are also Tickers, which we'll look at next.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers represent a single event in the future.
	// You tell it how long to wait, and it provides a channel that gets notified at that time.
	// This Timer will wait 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// This blocks on the timer's channel C until it sends a value showing that it fired.
	// For a simple wait, this is overcomplicated - could just use time.Sleep()
	<-timer1.C
	fmt.Println("Timer 1 fired")
	stop1 := timer1.Stop()
	if !stop1 {
		fmt.Println("Timer 1 was already stopped!")
	}

	// A Timer can be cancelled though!
	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
