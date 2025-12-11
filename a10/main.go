package main

import "fmt"

func main() {
	var nopositif uint8 = 89
	var nonegatif = -1243423644

	fmt.Printf("bilangan positif: %d\n", nopositif)
	fmt.Printf("bilangan negatif: %d\n", nonegatif)

	fmt.Printf("bilangan positif: %T\n", nopositif)
	fmt.Printf("bilangan negatif: %T\n", nonegatif)

	fmt.Println("bilangan positif:", nopositif)
	fmt.Println("bilangan negatif:", nonegatif)
}
