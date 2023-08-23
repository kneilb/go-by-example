// https://gobyexample.com/context
//
// In the previous example, we set up a simple HTTP server.
// HTTP servers are useful for showing how to use context.Context for controlling cancellation.
// A Context carries deadlines, cancellation signals, and other request-scoped values across API-boundaries & goroutines.
//
// To test:
//
// Run the server in the background
//   go run context-in-http-servers.go &
//
// Then simulate a client request to /hello, but hit Ctrl+C shortly afterward to signal cancellation.
//   curl localhost:8090/hello

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// The net/http machinery creates a context.Context for each request.
	// You can access it with the .Context() method.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait for a while before returning a response to the client.
	// This simulates the server doing "some work".
	// While working, peer at the context's Done() channel.
	// Cancel the work & return ASAP if signalled.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")

	case <-ctx.Done():

		// The context's Err() method returns an error that explains why the Done() channel was closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8092", nil)
}
