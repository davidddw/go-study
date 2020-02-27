package main

import (
	"fmt"
)

// if 判断
func main() {
	// 方式1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 方式2
	var j = 5
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	// 方式3
	var k = 5
	for k < 10 {
		fmt.Println(k)
		k++
	}

	// for-range 循环
	s := "hello world"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
