package main

import (
	"fmt"
)

func main() {
	s := "hello world"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d", i)
	}
}