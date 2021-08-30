package main

import (
	"fmt"
)

func main() {
	//declaring and initializing variable b
	var b string = "Hello World"

	//declaring string pointer, a
	a := new(string)
	//initialization of pointer
	a = &b
	fmt.Println("value of b", b)
	fmt.Println("address of b", &b)
	fmt.Println("value of a", a)
}
