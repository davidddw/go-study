package main

import "fmt"

// 结构体占用一块连续的内存空间
type x struct {
	a int8
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(100),
		b: int8(110),
		c: int8(120),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
}
