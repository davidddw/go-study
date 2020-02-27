package main

import "fmt"

//空接口

func assert(a interface{}) {
	fmt.Printf("type: %T\n", a)
	str, ok := a.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("错误类型")
	}
}

func assert2(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("x is a string %s\n", v)
	case int:
		fmt.Printf("x is a integer %d\n", v)
	case bool:
		fmt.Printf("x is a boolean %v\n", v)
	case float64:
		fmt.Printf("x is a float64 %v\n", v)
	default:
		fmt.Println("错误类型")
	}
}

func main() {
	assert(100)
	assert("100")

	assert2("100")
	assert2(true)
	assert2(1.23)
}
