package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var todos []string
	var pilihan int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=====/selamat datang di app todo/=====")
		fmt.Println("1. Create Todo")
		fmt.Println("2. Look Todo")
		fmt.Println("3. Edit Todo")
		fmt.Println("4. Hapus Todo")
		fmt.Println("5. Keluar")

		fmt.Print("Pilih menu (1-5): ")
		fmt.Scanln(&pilihan)

		switch pilihan {

		case 1:
			// Create Todo
			fmt.Print("Masukan Todo: ")
			todo, _ := reader.ReadString('\n')
			todo = strings.TrimSpace(todo)

			todos = append(todos, todo)
			fmt.Println("Todo Berhasil di tambahkan")

		case 2:
			// Look Todo
			fmt.Println("\nDaftar Todo")

			if len(todos) == 0 {
				fmt.Println("Todo masih kosong")
			} else {
				for i, todo := range todos {
					fmt.Printf("%d. %s\n", i+1, todo)
				}
			}

		case 3:
			// Edit Todo
			if len(todos) == 0 {
				fmt.Println("\nTidak ada Todo yang bisa di edit")
				break
			}

			fmt.Println("\nDaftar Todo: ")
			for i, todo := range todos {
				fmt.Printf("%d. %s\n", i+1, todo)
			}

			var nomor int
			fmt.Print("Edit Todo No: ")
			fmt.Scanln(&nomor)

			if nomor < 1 || nomor > len(todos) {
				fmt.Println("Nomor tidak valid!!")
			} else {
				fmt.Print("Masukan todo baru: ")
				todoBaru, _ := reader.ReadString('\n')
				todoBaru = strings.TrimSpace(todoBaru)

				todos[nomor-1] = todoBaru
				fmt.Println("Todo Berhasil di edit!!")
			}

		case 4:
			// Hapus Todo
			if len(todos) == 0 {
				fmt.Println("\nTidak ada Todo yang bisa di hapus")
				break
			}

			fmt.Println("\nDaftar Todo: ")
			for i, todo := range todos {
				fmt.Printf("%d. %s\n", i+1, todo)
			}

			var nomor int
			fmt.Print("Hapus Todo No: ")
			fmt.Scanln(&nomor)

			if nomor < 1 || nomor > len(todos) {
				fmt.Println("Nomor tidak valid!!")
			} else {
				todos = append(todos[:nomor-1], todos[nomor:]...)
				fmt.Println("🗑Todo Berhasil di hapus!!")
			}

		case 5:
			// Keluar
			fmt.Println("THANK UUU")
			return

		default:
			fmt.Println("Pilihan tidak valid!!")

		}
	}
}
