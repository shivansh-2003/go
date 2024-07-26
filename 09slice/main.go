package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to slice")
	var course = []string{"reactjs", "numpy", "swift", "python", "ruby"}
	fmt.Println(course)

	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)

	var course = []string{"reactjs", "numpy", "swift"}

}
