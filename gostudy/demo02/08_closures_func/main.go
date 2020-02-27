package main

import "fmt"

// 闭包

func adder100() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder100()
	ret2 := ret(200)
	fmt.Println(ret2)
}
