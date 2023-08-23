// https://gobyexample.com/struct-embedding

package main

import "fmt"

// A base struct, with some basic functionality
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// a container struct, which embeds base
type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{num: 1},
		str:  "my name is benny",
	}

	// Can access the base's fields directly, e.g. co.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Or spell out the full path
	fmt.Println("or num:", co.base.num)

	// We can also call the functions of base directly on the embedding type
	fmt.Println("describe:", co.describe())

	// We can use struct embedding to implement an interface, too
	// a container implements describer, because it embeds base
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
