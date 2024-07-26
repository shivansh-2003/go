package main

import (
	"fmt"
)

func main() {
	fmt.Println("map in go")
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["go"] = "golang"
	languages["py"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["js"])
	delete(languages, "py")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("for key %v,value is %v\n", key, value)
	}
}
