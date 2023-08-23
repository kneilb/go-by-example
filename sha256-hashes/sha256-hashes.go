// https://gobyexample.com/sha256-hashes
//
// SHA256 hashes are used in TLS/SSL certs.
// Here's how to compute them in Go.
//
// There are plenty of other hashes available in crypto/* packages.
// For use in real crypto, research hash strength!
// https://en.wikipedia.org/wiki/Cryptographic_hash_function

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Start a new hash.
	h := sha256.New()

	// Write expects bytes; coerce strings to byte slices.
	h.Write([]byte(s))

	// Get the finished result as a byte slice.
	// The argument can be used to append to an existing byte slice.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
