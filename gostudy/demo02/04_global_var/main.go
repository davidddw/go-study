package main

import "fmt"

// 变量作用域

var x = 100

func f1() {
	// 先在函数内部找
	// 在全局找
	fmt.Println(x)
}

func main() {
	f1()
}
