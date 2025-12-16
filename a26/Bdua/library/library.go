package library

import "fmt"

var Student = struct {
	Name string
	Age  int
}{}

func init() {
	Student.Name = "Hazel Clein Muhammad"
	Student.Age = 17

	fmt.Println("--> library.go/library.go.go imported")
}
