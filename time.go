// https://gobyexample.com/time
//
// Examples of time & duration

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Current time
	now := time.Now()
	p(now)

	// Build a time struct, providing year, month etc.
	// Always associated with a Location (time zone).
	// This is when John was born.
	then := time.Date(2012, 2, 10, 11, 41, 58, 123456789, time.UTC)
	p(then)

	// Extract all the bits
	p(then.Year())
	p(then.Month())
	p(then.Day())

	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())

	p(then.Location())

	// Determine what day of the week it was (John was born on a Friday).
	p(then.Weekday())

	// Compare two times
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub returns a Duration representing the interval between 2 times.
	diff := now.Sub(then)
	p(diff)

	// p(diff.Years())
	// p(diff.Days())

	// Compute the length in various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Use Add to advance or move backwards by a Duration.
	p(then.Add(diff))
	p(then.Add(-diff))
}
