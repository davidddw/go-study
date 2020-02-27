package main

import "fmt"

func f1() {
	fmt.Println("hello world")
}

func f2(name string) {
	fmt.Println("hello", name)
}

func f3(x int, y int) int {
	return x + y
}

func f4(x, y int) int {
	return x + y
}

func f5(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func f6(x, y int) (sum int) {
	sum = x + y
	return
}

func f7(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("world")
	fmt.Println(f3(100, 200))
	fmt.Println(f4(100, 200))
	fmt.Println(f5(1, 2, 3, 4, 5))
	fmt.Println(f6(1, 2))
	fmt.Println(f7(10, 4))
}
