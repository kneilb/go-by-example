// https://gobyexample.com/errors

package main

import (
	"errors"
	"fmt"
)

// By convention, error is the last return value
// They have the type "error"
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New() constructs a basic "error" value with the given message
		return -1, errors.New("42 is not acceptable, exploding")
	}

	return arg + 3, nil
}

// We can use custom types as Errors by implementing the Error() method on them
// (which is the only method in the error interface)
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// Test out the error-returning functions
	// The inline error check is good idiomatic Go
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// To programmatically use the data in a custom error,
	// get the error as an instance of the custom type via assertion
	// (which is like a cast in C)
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
