package main

import "fmt"

func main() {

	var age int = 14
	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a kid")
	}

	if age1 := 15; age1 >= 18 {
		fmt.Println("person is an adult", age1)
	} else if age1 >= 12 {
		fmt.Println("person is teenger", age1)
	}

	//  go does not have ternary operation
}
