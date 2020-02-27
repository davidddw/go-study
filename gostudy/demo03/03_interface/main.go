package main

import "fmt"

//接口的实例2

type car interface {
	run()
}

// 法拉利
type ferrari struct {
	brand string
}

// 保时捷
type porsche struct {
	brand string
}

func drive(c car) {
	c.run()
}

func (f ferrari) run() {
	fmt.Printf("%s速度200迈~\n", f.brand)
}

func (p porsche) run() {
	fmt.Printf("%s速度700迈~\n", p.brand)
}

func main() {
	var f1 = ferrari{
		brand: "法拉利",
	}
	var p1 = porsche{
		brand: "保时捷",
	}
	drive(f1)
	drive(p1)
}
