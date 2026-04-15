package main

import(
	"fmt"
)

var age int = 42
var name string = "James Bond"
var isActor bool = true

func main() {
	// Exercício 3
	s := fmt.Sprintf("%s, %d, %v", name, age, isActor)
	fmt.Println(s)
}