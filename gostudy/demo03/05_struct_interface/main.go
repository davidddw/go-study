package main

import "fmt"

//接口的实例2

type eater interface {
	eat(string)
}

type mover interface {
	move()
}

type animal interface {
	eater
	mover
}

type cat struct {
	name string
	feet int8
}

// 方法使用指针接受者
func (c *cat) move() {
	fmt.Println("走猫步")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s~\n", food)
}

func main() {
	var a1 animal
	c := &cat{"jerry", 4}
	a1 = c
	fmt.Println(a1)
}
