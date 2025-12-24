package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {

	fmt.Printf("%b\n", data.age) 

	fmt.Printf("%c\n", 1400) 
	fmt.Printf("%c\n", 1235) 

	fmt.Printf("%d\n", data.age) 

	fmt.Printf("%e\n", data.height) 
	fmt.Printf("%E\n", data.height) 

	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.9f\n", data.height) 
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)  

	fmt.Printf("%e\n", 0.123123123123) // 1.231231e-01
	fmt.Printf("%f\n", 0.123123123123) // 0.123123
	fmt.Printf("%g\n", 0.123123123123) // 0.123123123123

	fmt.Printf("%g\n", 0.12)   
	fmt.Printf("%.5g\n", 0.12) 

	fmt.Printf("%o\n", data.age) 

	fmt.Printf("%p\n", &data.name) 

	fmt.Printf("%q\n", `" name \ height "`) 

	fmt.Printf("%s\n", data.name) 

	fmt.Printf("%t\n", data.isGraduated) 

	fmt.Printf("%T\n", data.name)        // string
	fmt.Printf("%T\n", data.height)      // float64
	fmt.Printf("%T\n", data.age)         // int32
	fmt.Printf("%T\n", data.isGraduated) // bool
	fmt.Printf("%T\n", data.hobbies)     // []string

	fmt.Printf("%v\n", data)

	fmt.Printf("%+v\n", data)

	fmt.Printf("%#v\n", data)

	var people = struct {
		name   string
		height float64
	}{
		name:   "wick",
		height: 182.5,
	}
	var d = data.name

	fmt.Printf("%#v\n", people)


	fmt.Printf("%x\n", data.age)                     
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3]) 
	fmt.Printf("%x\n", d)                            
	fmt.Printf("%%\n")                               
}
