package main

import "flag"
import "fmt"

func main() {
	var data1 = flag.String("name", "anonymous", "type your name")
	fmt.Println(*data1)

	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	fmt.Println(data2)
}
