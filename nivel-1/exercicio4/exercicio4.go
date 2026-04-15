package main

import (
	"fmt"
)

type cat int
var x cat

func main() {
	// Exercício 4
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Print(x)
}