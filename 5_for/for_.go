package main

import "fmt"

// for is the only loop which construct go.
func main() {

	// while loop
	// i:=1
	// for i<=3 {
	// 	fmt.Println(i);
	// 	i = i+1;
	// }

	// infinite loop
	// for {
	// 	println("1")
	// }

	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println("curr i is ", i)
	}

	for i := range 10 {
		fmt.Println(i)
	}
}
