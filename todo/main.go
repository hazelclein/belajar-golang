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
		fmt.Println("3. Hapus Todo")
		fmt.Println("4. Keluar")

		fmt.Println("Pilih menu (1-4): ")
		fmt.Scanln(&pilihan)

		switch pilihan {

		case 1:
			fmt.Print("Masukan Todo: ")
			todo, _ := reader.ReadString('\n')
			todo = strings.TrimSpace(todo)

			todos = append(todos, todo)
			fmt.Println("Todo Berhasil di tambahkan")

		case 2:
			fmt.Println("\nDaftar Todo")

			if len(todos) == 0 {
				fmt.Println("Todo masih kosong")
			} else {
				for i, todo := range todos {
					fmt.Printf("%d. %s\n", i+1, todo)
				}
			}

		case 3:
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
				fmt.Println("Todo Berhasil di hapus!!")
			}

		case 4:
			fmt.Println("THANK UUU")
			return

		default:
			fmt.Println("Pilihan tidak valid!!")

		}
	}
}
