package main

import "fmt"

func main() {
	n := 18
	p := &n
	fmt.Printf("type %T, value %v, value %v\n", p, p, *p)

	// var a *int
	// *a = 100
	// fmt.Println(*a)

	var a *int = new(int)
	*a = 100
	fmt.Println(*a)
}
