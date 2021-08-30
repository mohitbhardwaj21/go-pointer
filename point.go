package main

import (
	"fmt"
)

func main() {
	var b string = "Hello World"
	var a *string = &b
	fmt.Println("value of b ->", b)
	fmt.Println("address of b ->", &b)
	fmt.Println("value of a ->", a)
	fmt.Println("accessing value of b using pointer a ->", *a)
}
