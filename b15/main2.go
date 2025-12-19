package main

import "fmt"
import "net/http"
import "encoding/json"

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"hazel clein", 24},
		{"m arievan", 23},
		{"m rasya", 22},
		{"doni", 21},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
