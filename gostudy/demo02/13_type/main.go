package main

import "fmt"

type myInt int    // 自定义类型
type myInt2 = int // 类型别名

func main() {
	var n myInt
	n = 100
	fmt.Printf("%T, %v\n", n, n)
	var m myInt2
	m = 100
	fmt.Printf("%T, %v\n", m, m)
	var c rune
	c = '中'
	fmt.Printf("%T, %c\n", c, c)
}
