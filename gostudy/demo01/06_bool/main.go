package main

import (
	"fmt"
)

func main() {
	b1 := true
	fmt.Printf("%T, value: %v\n", b1, b1)
	var b2 bool
	fmt.Printf("%T, value: %v\n", b2, b2)
}
