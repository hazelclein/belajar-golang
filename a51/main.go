package main 

import (
	"fmt"
	"net/http"
)

func index (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bagaimana kabarmu?")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo")
	})

	http.HandleFunc("/index", index)

	fmt.Println("server na leumpang ti port http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}