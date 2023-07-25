// https://gobyexample.com/channel-directions

package main

import "fmt"

// This is a send-only channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings is a receive-only channel
// pongs is a send-only channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "message for you, sir")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
