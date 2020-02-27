package main

import "fmt"

//接口的实例2

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 方法使用值接受者
// func (c cat) move() {
// 	fmt.Println("走猫步")
// }

// func (c cat) eat(food string) {
// 	fmt.Printf("猫吃%s~\n", food)
// }

// 方法使用指针接受者
func (c *cat) move() {
	fmt.Println("走猫步")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s~\n", food)
}

func main() {
	var a1 animal

	//c1 := cat{"tom", 4}
	c2 := &cat{"jerry", 4}

	// 使用值接收者实现接口， 接口体类型和结构体的指针都能存
	// 使用指针接收者实现的接口，只能存结构体的指针类型的变量
	//a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
