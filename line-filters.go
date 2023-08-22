// https://gobyexample.com/line-filters
//
// A line filter reads from stdin, processes input & prints the result on stdout.
// grep & sed are examples of line filters.
//
// This simple example capitalises all input.
//
// echo 'hello'   > /tmp/lines
// echo 'filter' >> /tmp/lines
// cat /tmp/lines | go run line-filters.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
