package main

import "fmt"

// fmt

func main() {
	var s string
	fmt.Print("输入字符串：")
	fmt.Scanln(&s)
	fmt.Println("用户输入的是：", s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Print("输入用户信息：")
	fmt.Scanf("%s %d %s", &name, &age, &class)
	fmt.Println(name, age, class)
}
