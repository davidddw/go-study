package main

import "fmt"

// 函数的定义
func sum(x int, y int) (ret int) {
	ret = x + y
	return
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数有返回值
func f3() int {
	return 3
}

// 命名返回值
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多个返回值
func f5() (int, string) {
	return 1, "ok"
}

// 参数类型简写
func f6(x, y int) int {
	return x + y
}

// 变长参数类型简写
func f7(x int, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	fmt.Println(sum(1, 2))
	_, n := f5()
	fmt.Println(n)
	f7(1)
	f7(1, 2, 3, 4)
}
