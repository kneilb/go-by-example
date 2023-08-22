// https://gobyexample.com/random-numbers
//
// math/rand does pseudorandom numbers.
// Use crypto/rand for "proper" ones.
//
// More details here: https://pkg.go.dev/math/rand

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// rand.Intn(arg) produces a random int n 0 <= n < arg
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0
	fmt.Println(rand.Float64())

	// This can be used to produce random floats in other ranges
	// e.g. 5.0 <= f' < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Apparently the default number generator is deterministic, and will produce the same sequence each time.
	// However, this doesn't seem to work in real life... Multiple runs give different results for the above.
	// Aaaaanyway...
	//
	// To produce varying sequences, you "have" to seed the RNG.
	// This still isn't good enough for crypto, obviously. See above.

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Call the resulting rand.Rand just like the functions in the rand package itself.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// If you seed a source with the same number, it produces the same sequence.
	// For real, this time.

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
}
