package main

import "fmt"

// 递归 阶乘
// 5! = 5*4*3*2*1

func taijie(n int64) int64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func f(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

func main() {
	fmt.Println(f(6))
	fmt.Println(taijie(6))
}
