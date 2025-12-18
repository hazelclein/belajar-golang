package main

import "encoding/json"
import "net/http"
import "fmt"

type student struct {
    ID     string  `json:"id"`
    Name   string  `json:"name"`
    Grade  int     `json:"grade"`
    Tinggi float64 `json:"tinggi"`
}

var data = []student{
    {"E001", "yanto", 21, 170.5},
    {"W001", "budi", 22, 175.2},
    {"B001", "nia", 23, 180.0},
    {"B002", "jono", 23, 178.5},
    {"B003", "epan", 23, 175.5},
    {"B004", "doni", 23, 175.5},
    {"C001", "odih", 24, 182.3},
    {"C002", "agus", 24, 180.5},
    {"C003", "julianto", 24, 180.5},
    {"C004", "santoso", 24, 179.0},
    {"C005", "bambang", 24, 180.5},
    {"C006", "dwi", 24, 180.5},
    {"C007", "tri", 24, 177.5},
}

func users(w http.ResponseWriter, r *http.Request) {
    fmt.Println("[LOG] Request to /users")
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("[LOG] Request to /user with id=%s\n", r.FormValue("id"))
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var id = r.FormValue("id")
        var result []byte
        var err error

        for _, each := range data {
            if each.ID == id {
                result, err = json.Marshal(each)

                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                w.Write(result)
                return
            }
        }

        http.Error(w, "User not found", http.StatusBadRequest)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func main() {
    http.HandleFunc("/users", users)
    http.HandleFunc("/user", user)

    fmt.Println("Server running at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}