package main

import "fmt"

func main() {
	var s []string
	fmt.Println("uninitialised:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("made 3 long:", s, s == nil, len(s), cap(s))
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "zero"
	s[1] = "one"
	s[2] = "two"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "three")
	fmt.Println("appended:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, "four", "five")
	fmt.Println("appended:", s, "len:", len(s), "cap:", cap(s))

	c := make([]string, len(s))
	elems := copy(c, s)
	fmt.Println("copied", elems, "elements:", c)

	l := s[2:5] // x >= 2 && x < 5, i.e. s[2], s[3], s[4]
	fmt.Println("slice s[2:5]:", l)

	l = s[:5] // x < 5; up to, but excluding s[5]
	fmt.Println("slice s[:5]:", l)

	l = s[2:] // x >= 2; from 2 onwards
	fmt.Println("slice s[2:]:", l)

	// single line declaration / initialisation
	t := []string{"goat", "cheese", "ninja"}
	fmt.Println("declaration:", t)

	// You can make 2D slice structures as well.
	// Unlike with arrays, the inner slices can have different sizes.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}
