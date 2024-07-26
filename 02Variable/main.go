package main

import "fmt"

const login string="cdiifi"
func main() {
	var username string = "shivansh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isloggedIn bool = false
	fmt.Println(isloggedIn)
	fmt.Printf("variable is of type: %T \n", isloggedIn)

	var small uint8 = 255
	fmt.Println(small)
	fmt.Printf("variable is of type: %T \n", small)

	var Small float32 = 255.5645
	fmt.Println(Small)
	fmt.Printf("variable is of type: %T \n", Small)

	var website = "shivansh.m2003@gmail.com"
	fmt.Println(website)

	nuofuser := 30000
	fmt.Println(nuofuser)
}
