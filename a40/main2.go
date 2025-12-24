package main

import "fmt"
import "time"

func main() {
	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())
}

// Output:
// year: 2024 month: June
