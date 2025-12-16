package main

import (
	. "Asatu/library"
	f "fmt"
)

func main() {
	var s1 = Student{"Hazel", 17}
	f.Println("Nama: ", s1.Name)
	f.Println("Umur: ", s1.Age)
}
