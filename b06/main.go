package main

import "net/http"
import "fmt"
import "html/template"

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Hazel Clein Muhammad",
			Gender:  "male",
			Hobbies: []string{"Novel Creator", "Game Development", "Web Development"},
			Info:    Info{"luna enterprises", "Bogor, cimahpar"},
		}
		
		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
