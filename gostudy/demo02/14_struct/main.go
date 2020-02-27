package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明person类型变量p
	var p person
	// 通过字段赋值
	p.name = "孙悟空"
	p.age = 100
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Printf("type:%T value:%v\n", p, p)
	var p2 person
	p2.name = "猪八戒"
	p2.age = 200
	fmt.Printf("type:%T value:%v\n", p2, p2)

	// 匿名结构体
	var s struct {
		x string
		y int
	}
	s.x = "hello"
	s.y = 2
	fmt.Printf("type:%T value:%v\n", s, s)
}
