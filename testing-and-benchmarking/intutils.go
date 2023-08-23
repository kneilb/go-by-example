// https://gobyexample.com/testing-and-benchmarking
//
// Unit testing is important, m'kay?
// The testing package provides tools to write them, and "go test" runs them.
//
// I've split the tests & implementation into separate files, to show how it would work in the "real world".
// I've also given this a more sane package name.
// The tests are in intutils_test.go

package intutils

// We're going to test this function, which implements a simple integer minimum.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
