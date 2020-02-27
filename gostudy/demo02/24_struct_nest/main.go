package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
	workPlace
}

type company struct {
	name    string
	address // 匿名嵌套结构体
}

func main() {
	p1 := person{
		name: "孙悟空",
		age:  1000,
		address: address{
			"吉林省",
			"吉林市",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.address.province)
	fmt.Println(p1.address.city)
}
