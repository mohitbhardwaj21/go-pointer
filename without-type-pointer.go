package main

import (
	"fmt"
)

func main() {
	//declaring and initializing variable b
	var b string = "Hello World"
	//declaring and initialization pointer, a
	var a = &b
	fmt.Println("value of b", b)
	fmt.Println("address of b", &b)
	fmt.Println("value of a", a)
}
