package main

import "fmt"

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["吉林"] = 200
	m1["成都"] = 100
	m1["北京"] = 400
	fmt.Println(m1)

	fmt.Println(m1["成都"]) //如果不存在 为零值
	if v, ok := m1["上海"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("没找到")
	}

	//遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	//遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	//删除
	delete(m1, "成都")
	fmt.Println(m1)
}
