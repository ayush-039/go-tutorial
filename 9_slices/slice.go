package main

import "fmt"

func main() {
	// unintialized slice is nill.
	var nums = make([]int, 2, 5)
	// nums := []int{}
	nums[0] = 1
	nums[1] = 2
	fmt.Println(cap(nums))
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 5)
	nums = append(nums, 5)
	nums = append(nums, 5)
	nums = append(nums, 5)
	nums = append(nums, 5)
	fmt.Println(cap(nums))

	fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))
}
