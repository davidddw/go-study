package main

import "fmt"

// 常量定义之后不允许改变
const pi = 3.1415926

// 批量声明
const (
	statusOK = 200
	notFound = 404
)

// 如果没有赋值，默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota常量计数器
const (
	m1 = iota
	m2
	m3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
	d5, d6 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)

	fmt.Println("m1: ", m1)
	fmt.Println("m2: ", m2)
	fmt.Println("m3: ", m3)

	fmt.Println("d1: ", d1)
	fmt.Println("d2: ", d2)
	fmt.Println("d3: ", d3)
	fmt.Println("d4: ", d4)
	fmt.Println("d5: ", d5)
	fmt.Println("d6: ", d6)
}
