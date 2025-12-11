package main

import "fmt"

func main() {
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	// ---------------------

	dst = []string{"potato", "potato", "potato"}
	src = []string{"watermelon", "pinnaple"}
	n = copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
