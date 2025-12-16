package main

import (
	"Elima/library"
	"fmt"
)

func main() {
	var s1 = library.Student{"Hazel", 21}
	fmt.Println("nama: ", s1.Name)
	fmt.Printf("umur :", s1.Age)
}
