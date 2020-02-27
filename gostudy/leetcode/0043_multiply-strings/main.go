package main

import (
	"fmt"
	"strings"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，
它们的乘积也表示为字符串形式。
输入: num1 = "123", num2 = "456"
输出: "56088"
*/

func main() {
	num1 := "213123213"
	num2 := "151231231231"
	target := multiply(num1, num2)
	fmt.Println(target)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1 := len(num1)
	len2 := len(num2)
	num := make([]byte, len1+len2)
	var numR []byte
	for i := len1 - 1; i >= 0; i-- {
		t := num1[i] - '0'
		for j := len2 - 1; j >= 0; j-- {
			sum := t * (num2[j] - '0')
			num[i+j+1] += sum % 10
			num[i+j] += (num[i+j+1]/10 + sum/10)
			num[i+j+1] %= 10
		}
	}
	for i := 0; i < len1+len2; i++ {
		numR = append(numR, num[i]+'0')
	}
	retStr := string(numR)
	return strings.TrimLeft(retStr, "0")
}
