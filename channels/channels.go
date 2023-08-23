// https://gobyexample.com/channels

package main

import "fmt"

func main() {
	// Create a new channel, which carries strings
	messages := make(chan string)

	// Send a value into the channel
	// We do this from an async goroutine
	go func() { messages <- "ping" }()

	// Trying to do this in the main thread also causes a deadlock
	// fatal error: all goroutines are asleep - deadlock!
	// messages <- "ping"

	// Read a value from the channel
	// This blocks the receiver, so if we didn't send anything above we get this:
	// fatal error: all goroutines are asleep - deadlock!
	msg := <-messages
	fmt.Println(msg)
}
