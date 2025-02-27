package main

import "fmt"

func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	name := []string{"Ayush", "barot", "brahmbhatt"}
	printSlice(name)
}
