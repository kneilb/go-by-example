package main

import "fmt"

// a rect struct
type rect struct {
	width, height int
}

// method with a "receiver type" of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// method with a "value receiver" (i.e. rect)
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	// Create the struct instance
	r := rect{width: 10, height: 5}

	// Call the struct methods
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Automatic conversion between values & pointers
	// Pointer vs value on method calls is the same as C++:
	// - values are copies, changes apply only to that copy
	// - pointers allow the original to be changed, no copying
	// - unfortunately there don't seem to be pointers to const...
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
