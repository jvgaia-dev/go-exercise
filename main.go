package main

import (
	"fmt"
)

type hotdog int
var b hotdog = 101

func main() {
	x := 32
	fmt.Println(x)

	x = int(b)
	fmt.Println(x)
}