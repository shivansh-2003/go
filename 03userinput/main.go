package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating of pizza ")

	//comma ok
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating,", input)
	fmt.Printf("Type of the rating is %T", input)

}
