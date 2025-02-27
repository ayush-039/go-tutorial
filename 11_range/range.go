package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	// sum := 0
	var sum int = 0
	for idx, num := range nums {
		fmt.Println(idx)
		sum = sum + num
	}
	fmt.Println(sum)

	// creating map
	m := map[string]string{"fname:": "Ayush", "lname:": "Barot"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// unicode code point rune
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
