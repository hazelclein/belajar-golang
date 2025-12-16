package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Produk struct {
	kode string
	nama string
	harga float64
	stok int
}

type ItemBelanja struct {
	produk  Produk
	Quantity int
	Subtotal float64
}

type Transaksi struct {
	NomorKota string
	Items	 []ItemBelanja
	TotalBelanja float64
	Bayar float64
	Kembalian float64
	Waktu time.Time
}

var daftarProduk = []Produk{
	{Kode : "P001", Nama: "Pensil", Harga: 2000, Stok: 100},
	{Kode : "P002", Nama: "Buku Tulis", Harga: 5000, Stok: 50},
	{Kode : "P003", Nama: "Penghapus", Harga: 1500, Stok: 75},
	{Kode : "P004", Nama: "Spidol", Harga: 3000, Stok: 60},
	{Kode : "P005", Nama: "Penggaris", Harga: 2500, Stok: 80},
	{Kode : "P006", Nama: "pulpen", Harga: 1800, Stok: 90},
	{Kode : "P007", Nama: "Buku Gambar", Harga: 7000, Stok: 40},
	{Kode : "P008", Nama: "Kertas HVS", Harga: 6000, Stok: 70},
	{Kode : "P009", Nama: "Stapler", Harga: 8000, Stok: 30},
	{Kode : "P010", Nama: "Lem Kertas", Harga: 3500, Stok: 85},
	{Kode : "P011", Nama: "Tipe-X", Harga: 4000, Stok: 55},
	{Kode : "P012", Nama: "Map Plastik", Harga: 4500, Stok: 65},
	{Kode : "P013", Nama: "Clip Kertas", Harga: 1200, Stok: 95},
	{Kode : "P014", Nama: "Penggaris Segitiga", Harga: 5500, Stok: 45},
	{Kode : "P015", Nama: "Kalkulator Saku", Harga: 15000, Stok: 25},
	{Kode : "P016", Nama: "Penggaris Panjang", Harga: 9000, Stok: 35},
	{Kode : "P017", Nama: "Buku Agenda", Harga: 8000, Stok: 50},
	{Kode : "P018", Nama: "Kertas Warna", Harga: 7000, Stok: 60},
	{Kode : "P019", Nama: "Bolpoin Warna-warni", Harga: 5000, Stok: 75},
	{Kode : "P020", Nama: "sticker", Harga: 2200, Stok: 85},

}

func main() {
	var keranjang []ItemBelanja

	fmt.Println("Selamat datang di Toko nia!")
	fmt.Println()

	for {

		tampilkanMenu()

		var pilihan string
		fmt.Println("\nPilihan Menu (1-4):")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			tampilkanDaftarProduk()
		case "2":
			keranjang = tambahProduk(keranjang)
		case "3":
			if len(keranjang) == 0 {
				fmt.Println("\nKeranjang belanja kosong.")
			} else {
				keranjang = tampilkanKeranjang(keranjang)
			}
		case "4":
			if len(keranjang) == 0 {
				fmt.Println("\nTidak ada transaksi yang dilakukan.")
			} else {
				prosesCheckout(keranjang)
				keranjang = []ItemBelanja{}
			}

		default:
			fmt.Println("\nPilihan tidak valid. Silakan coba lagi.")


		}
	}
}

func tampilkanMenu() {
	fmt.Println("\n========================================")
	fmt.Println("           MENU UTAMA")
	fmt.Println("========================================")
	fmt.Println("1. Tampilkan Daftar Produk")
	fmt.Println("2. Tambah Produk ke Keranjang")
	fmt.Println("3. Tampilkan Keranjang Belanja")
	fmt.Println("4. Checkout dan Proses Pembayaran")
	fmt.Println("==========================================")
}

func tampilkanDaftarProduk() {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                     DAFTAR PRODUK                              ║")
	fmt.Println("╠══════╦══════════════════════════╦═══════════╦═════════════════╣")
	fmt.Println("║ Kode ║ Nama Produk              ║   Harga   ║      Stok       ║")
	fmt.Println("╠══════╬══════════════════════════╬═══════════╬═════════════════╣")

	for _, p := range daftarProduk {
		fmt.Printf("║ %-4s ║ %-24s ║ Rp %-6.0f ║ %3d ║\n", p.kode, p.nama, p.harga, p.stok)
	}

	fmt.Println("╚══════╩══════════════════════════╩═══════════╩═════════════════╝")
}

func cariProduk(kode string) *Produk {
	for i, p := range daftarProduk {
		if strings.ToUpper(daftarProduk[i].kode) == kode {
			return &daftarProduk[i]
		}
	}
	return nil
}

func tambahProduk(keranjang []ItemBelanja) []ItemBelanja {
	var kode string
	var qtystr string
}