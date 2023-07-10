package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Empty array: ", a)

	a[4] = 100

	fmt.Println("After set: ", a)
	fmt.Println("Get 4: ", a[4])
	fmt.Println("Length: ", len(a))

	// Don't _have_ to specify the dimensions (but maybe it's a slice then??)
	// b := []int{1, 2, 3, 4, 5}

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("inline declaration: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + 2*j
		}
	}
	fmt.Println("2D: ", twoD)
}
