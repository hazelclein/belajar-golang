package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	fmt.Println(fruits)
	fmt.Println(fruits[0:2])
	fmt.Println(fruits[0:4])
	fmt.Println(fruits[0:0])
	fmt.Println(fruits[4:4])
	fmt.Println(fruits[:])
	fmt.Println(fruits[2:])
	fmt.Println(fruits[:2])
}
