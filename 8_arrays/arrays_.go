package main

import "fmt"

// numbered sequence of specific length.
func main() {
	// default value
	// int -> 0 string->"" bool-> false
	// var nums [4]int
	// for i := 0; i < len(nums); i++ {
	// 	nums[i] = 2*i + 1
	// }
	// fmt.Println(len(nums))
	// // print whole array
	// fmt.Println(nums)
	// // iterating over array.
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// // want to add elements while declaring array

	// nums1 := [4]int{1, 2, 3, 4}
	// fmt.Println(nums1)

	// 2d array.

	nums2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums2)

}
