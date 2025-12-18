package main

import "net/http"
import "fmt"
import "html/template"

type DataDiri struct {
	Name    string
	Alias   string
	Keahlian []string
}

func (s DataDiri) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = DataDiri{
			Name:    "Hazel Clein Muhammad",
			Alias:   "Oce",
			Keahlian: []string{"Game Development", "Web Development", "Author Novel"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at http://localhost:9000/view")
	http.ListenAndServe(":9000", nil)
}
