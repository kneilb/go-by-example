// https://gobyexample.com/testing-and-benchmarking
//
// Unit testing is important, m'kay?
// The testing package provides tools to write them, and "go test" runs them.
//
// These are the tests.
// The implementation we're testing is in intutils.go
//
// Run the tests using "go test -v"
// Run the benchmarks using "go test -bench=."

package intutils

import (
	"fmt"
	"testing"
)

// Create a test by writing a function whose name begins with Test.
// t.Error* will report failures, but continue execution.
// t.Fatal* will report failures and stop the test immediately.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Writing tests can be repetitive, so we can use this table-driven style.
// Inputs & expected outputs are listed in a table, and a loop walks over them.
//
// t.Run enables running "subtests", one per table entry.
// These are shown separately when running with "go test -v"
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark tests typically live in _test.go files, and are named starting with Benchmark.
// The testing runner executes each function several times, increasing b.N until it gets a precise measurement.
//
// Typically the benchmark runs the function in a loop b.N times.
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
