package main

import "fmt"
import "time"

func main() {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())

}

// Output:
// 2015-09-02 08:04:00 	-> 2015-09-02 08:04:00 +0000 UTC
// 02/09/2015 WIB 		-> 2015-09-02 00:00:00 +0000 UTC
