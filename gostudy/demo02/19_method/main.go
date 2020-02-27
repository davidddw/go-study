package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func newPerson(name string, age int, gender string) *person {
	return &person{
		name:   name,
		age:    age,
		gender: gender,
	}
}

// Fight 方法是作用于特定类型的函数
// 多用类型名首字母小写
func (p *person) Fight() {
	fmt.Printf("%s fighting!\n", p.name)
}

// Ageing 变老
// 使用指针接收
// 1. 需要修改接收者的值
// 2. 接受者是拷贝代价大的对象
// 3. 保证一致性
func (p *person) Ageing() {
	p.age++
}

func (p *person) Dream() {
	fmt.Println("dream")
}

// 使用值接收
func (p person) ageing() {
	p.age++
}

func main() {
	p1 := newPerson("孙悟空", 2000, "男")
	p1.Fight()
	fmt.Println(p1.age)
	p1.ageing()
	fmt.Println(p1.age)
	p1.Ageing()
	p1.Dream()
	fmt.Println(p1.age)
}
