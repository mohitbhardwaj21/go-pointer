package main

import (
	"fmt"
)

func main() {
	var c *string
	fmt.Println("value of c ->", c)
	fmt.Println("dereferencing pointer c ->", *c)
}
