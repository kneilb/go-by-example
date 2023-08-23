// https://gobyexample.com/non-blocking-channel-operations
//
// select with a default clause implements a _non-blocking_ send, receive or multi-way select

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receive
	// if a value is available, it will be received (first case)
	// if not, we'll fail to receive (default)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// non-blocking send
	// if there is a buffer or a receiver, the value can be sent (first case)
	// if not, we fail to send the message (default)
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// multi-way non-blocking select
	// attempt non-blocking receives on both messages & signals
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
