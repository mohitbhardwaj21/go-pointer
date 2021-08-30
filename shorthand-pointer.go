package main

import (
	"fmt"
)

func main() {
	//declaring and initializing variable b
	b := 20
	//declaring and initialization pointer, a
	a := &b
	fmt.Println("value of b", b)
	fmt.Println("address of b", &b)
	fmt.Println("value of a", a)
}
