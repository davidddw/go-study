package main

import "fmt"

type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func (m myInt) String() string {
	return fmt.Sprintf("我的值是: %d", m)
}

func main() {
	//方法1
	//var x int32
	//x = 100
	//方法2
	//var x int32 = 100
	//方法3
	//var x = int32(100)
	//方法4
	// x := int32(100)
	// fmt.Println(x)

	// var m myInt
	// m = 100
	// fmt.Println(m)
}
