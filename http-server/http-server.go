// https://gobyexample.com/http-server
//
// It's easy to write a basic HTTP server using the net/http package.
// More docs here: https://pkg.go.dev/net/http#hdr-Servers
//
// $ go run http-servers.go &
// Access the /hello route.
//
// $ curl localhost:8092/hello
// hello
//
// curl localhost:8092/headers
// User-Agent, curl/7.68.0
// Accept, */*

package main

import (
	"fmt"
	"net/http"
)

// Handlers are fundamental to net/http servers.
// A handler is an object that implements the http.Handler interface.
// A common way to write one is by using the http.HandlerFunc adapter on functions with the right signature.
//
// Functions acting as handlers take an http.ResponseWriter and http.Request as arguments.
// The ResponseWriter is used to fill in the HTTP response.
func hello(w http.ResponseWriter, req *http.Request) {

	// Our simple response.
	fmt.Fprintf(w, "hello\n")
}

// This handler is a bit more sophisticated; it reads all the HTTP request headers and echoes them in the response body.
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v, %v\n", name, h)
		}
	}
}

func main() {

	// Register our handlers on server routes using the http.HandleFunc convenience function.
	// It sets up the default router in the net/http package, and takes a function as an argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Finally call ListenAndServe with the port and a handler.
	// nil tells it to use the default router we've just set up.
	err := http.ListenAndServe(":8092", nil)
	if err != nil {
		panic(err)
	}
}
