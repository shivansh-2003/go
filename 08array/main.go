package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcometo array in golang ")
	var fruit [4]string
	fruit[0] = "apple"
	fruit[1] = "orange"
	fruit[2] = "honey"
	fmt.Println(fruit)
	fmt.Println(len(fruit))
}
