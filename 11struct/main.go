package main

import (
	"fmt"
)

func main() {
	fmt.Println("struct in golang")
	shivansh := User{"shivansh", "shivansh.m2003@gamil.com", true, 23}
	fmt.Println(shivansh)
	fmt.Printf("%+v\n", shivansh)
	fmt.Printf("name is %v and email %v.", shivansh.Name, shivansh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
