package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(x, y int) int {
	return x + y
}

func f5(x func() int) func(int, int) int {
	return ff
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f3
	fmt.Printf("%T\n", b)
	f3(f2)
}
