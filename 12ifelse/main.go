package main

import (
	"fmt"
)

func main() {
	fmt.Println("if else in go")
	logic := 26
	var result string
	if logic < 10 {
		result = "regular user"
	} else if logic > 10 {
		result = "admin user"
	} else {
		result = "unknown user"
	}
	fmt.Println(result)
}
