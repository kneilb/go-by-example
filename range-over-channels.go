// https://gobyexample.com/range-over-channels
//
// We've seen before how for & range allow us to iterate over basic data structures.
// We can also use this to iterate over values received from a channel.

package main

import "fmt"

func main() {

	// Put 2 values in the queue, then close it.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it is received from queue.
	// Because we closed the channel above, the iteration terminates after receiving the elements.
	// This example also shows us getting the remaining values after the channel is closed.
	//
	// If we don't close the channel, we get the dreaded:
	// fatal error: all goroutines are asleep - deadlock!
	for elem := range queue {
		fmt.Println(elem)
	}
}
