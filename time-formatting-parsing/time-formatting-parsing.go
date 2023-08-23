// https://gobyexample.com/time-formatting-parsing
//
// Go supports time formatting & parsing via pattern-based layouts.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Format a time according to RFC3339, using the layout constant.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Parsing uses the same layout values.
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-02-10T11:41:58+00:00")
	p(t1)

	// Format and Parse use example-based layouts.
	// Usually, you'll use a constant from time, but you can also supply custom ones.
	// Layouts must use the reference time "Mon Jan 2 15:04:05 MST 2006" to show the pattern.
	// The reference time must be exactly as shown; the year 2006, 15 for the hour etc.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// For purely numeric representations, you can also use standard string formatting.
	// But why would you?
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse will return an error on malformed input, explaining the parsing problem.
	// This returns:
	// parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}
