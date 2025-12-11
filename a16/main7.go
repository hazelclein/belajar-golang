package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits))
	fmt.Println(len(bFruits))

	fmt.Println(fruits)
	fmt.Println(bFruits)

	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
}
