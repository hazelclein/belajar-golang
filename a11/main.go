package main

import "fmt"

func main() {
	const namaawal string = "hazel"
	const namaakhir = "clein"

	fmt.Printf("nama saya adalah %s %s\n", namaawal, namaakhir)
	fmt.Println("nama gua", namaawal, namaakhir)

	const minbuku int = 0
	const maxbuku int = 10
	const hargabuku int = 20000

	bukukebeli := 0

	if bukukebeli > maxbuku {
		fmt.Printf("pembelian melewati batas %d\n", maxbuku)
	} else if bukukebeli > minbuku && bukukebeli < maxbuku {
		totalharga := bukukebeli * hargabuku
		fmt.Printf("jumlah buku yang di beli %d\n", bukukebeli)
		fmt.Printf("total harga nya %d\n", totalharga)
	} else {
		fmt.Printf("tidak ada buku yang dibeli\n")
	}
	
}
