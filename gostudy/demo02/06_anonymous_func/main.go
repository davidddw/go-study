package main

import "fmt"

// 匿名函数

var f1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	f1(10, 20)

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
