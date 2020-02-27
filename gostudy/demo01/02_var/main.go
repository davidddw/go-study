package main

import "fmt"

// Go推荐驼峰格式命名

var (
	name string
	age  int
	isOK bool
)

func main() {
	name = "刘备"
	age = 42
	isOK = true
	// var hello string 声明不使用编译不通过

	fmt.Print(age)                 // 不会自动换行
	fmt.Println(isOK)              //自动加换行符
	fmt.Printf("name: %s\n", name) //%s占位符

	// 声明变量的同时赋值
	var s1 string = "关羽"
	fmt.Println(s1)

	// 类型推导
	var s2 = "张飞"
	fmt.Println(s2)

	// 简短变量声明
	s3 := "赵云"
	fmt.Println(s3)
}
