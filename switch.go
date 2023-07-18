// https://gobyexample.com/switch

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("It's one?")
	case 2:
		fmt.Println("It's a two.")
	case 3:
		fmt.Println("It appears to be a three.")
	}

	// Switch on a function call
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("I don't like Mondays")
	case time.Tuesday:
	case time.Thursday:
		fmt.Println("DHC Ride!")
	case time.Wednesday:
		fmt.Println("Just meh.")
	case time.Friday:
		fmt.Println("Hurray, it's Friday!")
	case time.Saturday:
	case time.Sunday:
		fmt.Println("It's the weekend!")
	}

	// Switch "on nothing"
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	default:
		fmt.Println("Afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Printf("Unknown type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(false)
	whatAmI(1)
	whatAmI("fishface")
	whatAmI(whatAmI)
	whatAmI([5]int{1, 2, 3, 4, 5})
	whatAmI([]string{"goat", "cheese", "ninja"})
	whatAmI(make([]string, 5))
}
