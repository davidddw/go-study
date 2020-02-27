package main

import "fmt"

func main() {
	var a int
	a = 100
	b := &a
	fmt.Printf("type:%T value:%v\n", a, a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("type:%T value:%v\n", b, b)
	fmt.Printf("type:%T value:%v\n", &b, &b)
}
