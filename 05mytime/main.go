package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of go lang")

	present := time.Now()
	fmt.Println(present)
	fmt.Println(present.Format("01-02-2006 Monday"))

	present = time.Now()
	fmt.Println(present)

}
