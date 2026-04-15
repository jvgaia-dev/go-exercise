package main

import (
	"fmt"
)

type cat int
var x cat 
var y int

func main() {
	// Exercício 5
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}