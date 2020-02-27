package main

import "fmt"

// 结构体嵌套

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type dog struct {
	feet   uint8
	animal // animal 有的方法 dog也有了
}

func (d dog) barking() {
	fmt.Printf("%s在叫 汪汪汪\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{name: "旺财"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.barking()
	d1.move()
}
