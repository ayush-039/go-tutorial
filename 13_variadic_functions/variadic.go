package main

import "fmt"

func sum(nums ...int) int {
	var sum int = 0
	for _, num := range nums {
		sum = sum + num
	}
	return sum
}

func main() {
	nums := []int{3, 4, 5, 6, 7, 9}
	result := sum(nums...)
	fmt.Println(result)
}
