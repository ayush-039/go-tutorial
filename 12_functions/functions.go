package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

func getLanguage() (string, string, string) {
	return "golang", "javascript", "c++"
}

func proccessIt(fn func(a int) int) {

}

func processIt2() func(a int) int {

	fn1 := func(a int) int {
		return 2 * a
	}

	return fn1
}

func main() {

	// result := add(3, 6)
	// fmt.Println(result)

	lang1, lang2, _ := getLanguage()
	fmt.Println(lang1, lang2)

}
