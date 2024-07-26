package main

import "fmt"

func main() {
	fmt.Println("welcometo a class on pointers")
	//var ptr *int
	//fmt.Println("value of pointer", ptr)

	num := 23
	var ptr = &num
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr + 3
	fmt.Println(*ptr)

}
