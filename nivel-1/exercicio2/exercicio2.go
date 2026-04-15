package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	// Exercício 2
	fmt.Printf("%v , %T\n", x, x)
	fmt.Printf("%v , %T\n", y, y)
	fmt.Printf("%v , %T", z, z)
}

//quando nao atribui valor, ele se torna zero