package main

import "fmt"

// 运算符
func main() {
	var (
		a = 5
		b = 2
	)
	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	b--
	fmt.Println(a, b)

	//关系运算符 相同类型
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	//逻辑运算符
	if a > 4 && b < 4 {
		fmt.Println("真")
	} else {
		fmt.Println("假")
	}
	if a < 4 || b < 4 {
		fmt.Println("真")
	} else {
		fmt.Println("假")
	}
	if !(a < 4) {
		fmt.Println("真")
	}

	//位运算符
	fmt.Println(a & b)  // 按位与
	fmt.Println(a | b)  // 按位或
	fmt.Println(a ^ b)  // 按位异或
	fmt.Println(a ^ b)  // 按位异或
	fmt.Println(a << 1) // 左移
	fmt.Println(a >> 1) // 由移

	//赋值运算符
	a += b
	fmt.Println(a, b)
	a -= b
	fmt.Println(a, b)
	a *= b
	fmt.Println(a, b)
	a /= b
	fmt.Println(a, b)
	a %= b
	fmt.Println(a, b)

	a >>= b
	fmt.Println(a, b)
}
