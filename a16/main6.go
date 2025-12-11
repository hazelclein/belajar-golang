package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(cFruits)
}
