package main

import "fmt"

func main() {
	var value = (((2+6)%3)*4 - 2) / 3
	var bener = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, bener)
}
