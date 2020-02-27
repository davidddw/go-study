package main

import (
	"fmt"

	"math/big"
)

func main() {
	// 初始化两个变量: a = 1, b = 2
	a := big.NewInt(1)
	b := big.NewInt(2)

	// 打印交换前的数值
	fmt.Printf("a = %v   b = %v\n", a, b)

	// 使用中间变量法进行交换
	tmp := big.NewInt(0)
	tmp.Set(a)
	a.Set(b)
	b.Set(tmp)
	// tmp := a
	// a, b = b, a
	// a = b
	// b = tmp

	// 交换完成, 对中间变量加100
	tmp.Add(tmp, big.NewInt(100))

	// 打印交换后的结果
	fmt.Printf("a = %v    b = %v   tmp = %v\n", a, b, tmp)
}
