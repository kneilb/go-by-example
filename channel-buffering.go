// https://gobyexample.com/channel-buffering

package main

import "fmt"

func main() {
	// A string channel that can buffer up to 2 values
	// This is different to the default channel, which will only accept a send if there is a corresponding receive ready
	messages := make(chan string, 2)

	// Because it's buffered, we can send values into it without the concurrent read
	messages <- "buffered"
	messages <- "channel"

	// Attempting a 3rd send gives a deadlock
	// messages <- "smells"

	// Later on, we can receive the values as normal
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Attempting a 3rd receive also gives a deadlock
	// fmt.Println(<-messages)
}
