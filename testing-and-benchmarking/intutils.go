// https://gobyexample.com/testing-and-benchmarking
//
// Unit testing is important, m'kay?
// The testing package provides tools to write them, and "go test" runs them.
//
// This code is in the "main" package, but could be anywhere.
// Test code typically lives in the same package as the code it tests.

package intutils

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
