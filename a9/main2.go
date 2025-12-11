package main

import "fmt"

func main() {
	nama_awal := "hazel"
	nama_akhir := "clein"

	fmt.Printf("halo hazel clein\n")
	fmt.Printf("halo %s %s\n", nama_awal, nama_akhir)
	fmt.Println("halo", nama_awal, nama_akhir+"!")
}
