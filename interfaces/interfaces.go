// https://gobyexample.com/interfaces

package main

import (
	"fmt"
	"math"
)

// define an interface for geometric shapes
type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type fish struct {
	length float64
	weight float64
}

// To implement an interface, just implement all the functions
// This is the implementation of geometry for rect
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// Implementation of geometry for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// generic measure function, which will work on any "geometry"
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	// both types implement geometry, so we can pass instances of the structs to measure()
	r := rect{width: 3, height: 4}
	measure(r)

	c := circle{radius: 5}
	measure(c)

	// If you try passing a type that doesn't implement the interface, something like this happens:
	// ./interfaces.go:62:10: cannot use f (variable of type fish) as geometry value in argument to measure: fish does not implement geometry (missing method area)
	// f := fish{length: 30, weight: 10000000}
	// measure(f)
}
