package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("caca cantik!"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("wleeee"))
	})

	fmt.Println("server na ngges lemmpang didieu localhost:9000")
	http.ListenAndServe(":9000", nil)
}