package main

import (
	"fmt"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
输入: 123
输出: 321
*/

func main() {
	num1 := 15342312
	fmt.Println(reverse(num1))
}

func reverse(x int) int {
	const max int = 0x7fffffff
	var b int
	for ; x != 0; b, x = b*10+x%10, x/10 {
	}
	if b > max || b < (-1*(max+1)) {
		return 0
	}
	return b
}
