// https://gobyexample.com/url-parsing
//
// URL parsing in Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// An example URL, with scheme, auth info, host, port, path, query params & query fragment
	s := "postgres://user:password@host.com:5432/path?k=v#f"

	// Parse the URL, ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Get the scheme
	fmt.Println(u.Scheme)

	// User contains the auth info; call Username & Password to get the individual values.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host contains the hostname & port, if present.
	// Use SplitHostPort to parse them.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Extract the path & fragment (bit after the #)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// RawQuery obtains query params in k=v format.
	// ParseQuery returns a map of strings to _slices of strings_.
	// Therefore need [0] to get the first value.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"])
	fmt.Println(m["k"][0])
}
