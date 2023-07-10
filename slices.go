package main

import "fmt"

func main() {
	var s []string
	fmt.Println("uninitialised:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("made 3 long:", s, s == nil, len(s), cap(s))
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "one"
	s[1] = "two"
	s[2] = "three"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "four")
	fmt.Println("appended:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, "five", "six")
	fmt.Println("appended:", s, "len: ", len(s), "cap:", cap(s))

	c := make([]string, len(s))
	elems := copy(c, s)
	fmt.Println("copied", elems, "elements:", c)

	fmt.Println([]string{"goat", "cheese", "ninja"})
}
