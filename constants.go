package main

import (
	"fmt"
	"math"
)

const s = "constant"

func main() {
	fmt.Println(s)

	// can't do this, it's a const
	// s += "fish"

	// const numbers don't have a type in go
	// You define it by casting them or using them
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
