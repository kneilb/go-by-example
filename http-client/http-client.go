// https://gobyexample.com/http-client
//
// The Go standard library comes with support for HTTP clients & servers in the net/http package.

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Issue an HTTP GET request to a server.
	// http.Get is a shortcut for creating an http.Client object and calling its Get method.
	// It uses the http.DefaultClient object, which has useful default settings.
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
