// https://gobyexample.com/signals
//
// Signal handling.
// For example, we might want a server to gracefully shut down when it receives a SIGTERM,
// or a command line tool to stop processing input if it receives a SIGINT.
// Here's how to do it in Go with channels.
//
// When executed, the program sits and waits for a SIGTERM or SIGINT (ctrl-C on the keyboard).

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go signal notifications send os.Signal values on a channel.
	// We'll create a channel to receive them; this channel should be buffered.
	sigs := make(chan os.Signal, 1)

	// signal.Notify registers the given channel to receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// We could receive sigs here in the main function, but let's do it in a separate goroutine, for more realism.
	done := make(chan bool, 1)

	// This goroutine does a blocking read for signals.
	// When it gets one, it'll print and notify the program that it can finish.
	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// The program will wait here until the goroutine above receives its signal and sends the value on done in response.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
