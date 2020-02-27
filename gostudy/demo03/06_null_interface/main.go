package main

import "fmt"

//空接口

func show(a interface{}) {
	fmt.Printf("type: %T value:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "张飒"
	m1["age"] = 20
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "跑"}
	fmt.Println(m1)

	show(false)
	show(m1)
	show(nil)
}
