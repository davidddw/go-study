package main

import "fmt"

//接口的实例

type speaker interface {
	speak()
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪")
}

func (p person) speak() {
	fmt.Println("啊啊")
}

func hit(s speaker) {
	s.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	hit(c1)
	hit(d1)
	hit(p1)
}
