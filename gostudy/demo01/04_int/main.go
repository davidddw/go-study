package main

import "fmt"

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1) //转成八进制
	fmt.Printf("%x\n", i1) //转成十六进制
	fmt.Printf("%b\n", i1) //转成二进制
	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制
	i3 := 0x12345ff
	fmt.Printf("%d\n", i3)
	//查看变量类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型
	i4 := int8(6)
	fmt.Printf("%T\n", i4)
}
