package main

import "fmt"

// 闭包

func outer(name string) func() {
	text := "Modified " + name
	foo := func() {
		fmt.Println(text)
	}
	return foo
}

func f1(x, y int) {
	fmt.Println(x + y)
}

func f2(x func(int, int), m int, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}

func main() {
	foo := outer("hello")
	foo()
	f2(f1, 500, 200)()
}
