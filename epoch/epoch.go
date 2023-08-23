// https://gobyexample.com/epoch
//
// The Unix Epoch:  https://en.wikipedia.org/wiki/Unix_time

package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)

	// Use time.Now() with Unix(), UnixMilli() or UnixNano() to get seconds, ms or ns since the Unix epoch.
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Convert integer seconds or nanoseconds since the Unix epoch into a time.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
