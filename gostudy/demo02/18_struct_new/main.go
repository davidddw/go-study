package main

import "fmt"

// 构造函数

type person struct {
	name   string
	age    int
	gender string
}

// 使用new开头
func newPerson(name string, age int, gender string) *person {
	return &person{
		name:   name,
		age:    age,
		gender: gender,
	}
}

func main() {
	p1 := newPerson("孙悟空", 2000, "男")
	p2 := newPerson("猪八戒", 3000, "男")
	p3 := newPerson("沙悟净", 1000, "男")
	fmt.Println(p1, p2, p3)
}
