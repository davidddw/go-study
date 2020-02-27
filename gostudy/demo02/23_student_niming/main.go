package main

import "fmt"

// 匿名结构体

type student struct {
	string
	int
}

func main() {
	var p1 = student{
		"孙悟空",
		1000,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
