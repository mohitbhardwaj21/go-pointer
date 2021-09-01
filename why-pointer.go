package main

import (
	"fmt"
)

func heroName(x string) {
	x = "Batman"
}

func main() {
	x := "Bruce Wayne"
	fmt.Println("value of x before function call ->", x)
	heroName(x)
	fmt.Println("value of x after function call ->", x)
}
