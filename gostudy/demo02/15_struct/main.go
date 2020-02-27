package main

import "fmt"

// 结构体指针

type person struct {
	name   string
	age    int
	gender string
}

func f(x person) {
	x.gender = "女" // 修改的是副本的值
}

func fp(x *person) {
	x.gender = "女" // 根据内存地址找到原变量，修改原来的值
}

func main() {
	var p1 person
	p1.name = "猪八戒"
	p1.age = 200
	p1.gender = "男"
	fmt.Printf("type:%T value:%v\n", p1, p1)
	f(p1)
	fmt.Printf("type:%T value:%v\n", p1, p1)
	fp(&p1)
	fmt.Printf("type:%T value:%v\n", p1, p1)

	//
	var p2 = new(person)
	p2.name = "golang"
	(*p2).gender = "男"
	fmt.Printf("type:%T value:%v\n", p2, p2)
	fmt.Printf("address:%p address:%v\n", p2, &p2)

	//  key value初始化
	var p3 = &person{
		name:   "孙悟空",
		age:    1000,
		gender: "男",
	}
	fmt.Printf("type:%T value:%v\n", p3, p3)

	// 值列表方式初始化， 顺序要一致
	p4 := &person{
		"沙悟净",
		1000,
		"男",
	}
	fmt.Printf("type:%T value:%v\n", p4, p4)

}
