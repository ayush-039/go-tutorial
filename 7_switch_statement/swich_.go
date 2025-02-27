package main

import (
	"fmt"
)

func main() {

	// simple switch
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend!")
	// default:
	// 	fmt.Println("it's workday!")
	// }

	// type Switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's an interger")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("golang")
	whoAmI(10.0)
	whoAmI(false)
}
