package main

import "fmt"

type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func (m myInt) String() string {
	return fmt.Sprintf("我的值是: %d", m)
}

func main() {
	m := myInt(100)
	n := int64(100)
	m.hello()
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(myInt(n))
}
