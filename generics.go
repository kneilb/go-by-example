// https://gobyexample.com/generics

package main

import "fmt"

// A generic function example
// Takes a map of "comparable" Keys, K to "any" Value V.
// "comparable" constraint means they can be compared with == and !=
// "any" is an alias for interface{}
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// A generic type example
// List is a singly-linked list with values of "any" type
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// It's also possible to define methods on generic types
// This is pretty much the same as doing it for regular types,
// but we have to keep the type parameters.
// The type is List[T], not just List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// Rely on "type inference"!
	// We don't have to specify what K & V are, the compiler just works them out
	fmt.Println("keys:", MapKeys(m))

	// Although we can do if we want...
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
