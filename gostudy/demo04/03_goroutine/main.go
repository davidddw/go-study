package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

func main() {
	// go hello(1)  // 开启一个goroutine执行hello函数
	// fmt.Println("main")
	for i := 0; i < 1000; i++ {
		//go hello(i)
		go func(i int) {
			fmt.Println("hello", i) // 用的是函数的参数i
		}(i)
	}
	fmt.Println("main")
	// main结束了，所有的goroutine都结束了
	time.Sleep(time.Second)
}
