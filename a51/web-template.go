package main

import (
	"html/template"
	"net/http"
	"fmt"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"nama": "Hazel Clein Muhammad",
			"Pesan": "Semoga harimu menyenangkan", 
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("server na leumpang ti port http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}