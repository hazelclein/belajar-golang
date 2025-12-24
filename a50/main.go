package main

import "fmt"
import "os"

var path = "C:/Users/hazel/Documents/hasilfileGO/test.php"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// cek file udah ada atau belum
	var _, err = os.Stat(path)

	// buat file baru kalo belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func main() {
	createFile()
}

//bikin file