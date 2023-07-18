// https://gobyexample.com/structs

package main

import "fmt"

// A struct with name & age fields
type person struct {
	name string
	age  int
}

// Construct a person struct with the passed name
// Returning the pointer is safe - the local variable survives...
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// Create a new struct instance
	fmt.Println(person{"Bob", 20})

	// Create a new struct, naming the fields
	fmt.Println(person{name: "Alice", age: 30})

	// Unspecified fields are zero-initialised
	// You have to name the fields to do this!
	fmt.Println(person{name: "Fred"})

	// Using & gets a pointer to the struct
	fmt.Println(&person{"Alex", 40})

	// Idiomatic creation via function - factory
	fmt.Println(newPerson("Sally"))

	// Access struct fields using a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Access pointer fields with a dot, too - pointer is automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs can be mutated, through pointer or directly
	sp.age = 51
	fmt.Println(s.age)

	s.age = 12
	fmt.Println(s.age)

	// anonymous structs are also possible
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
