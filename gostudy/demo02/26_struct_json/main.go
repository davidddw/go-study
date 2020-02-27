package main

import (
	"encoding/json"
	"fmt"
)

// 结构体json

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	d1 := person{
		Name: "孙悟空",
		Age:  1000,
	}
	p, err := json.Marshal(d1)
	if err != nil {
		fmt.Printf("marshal err: %v", err)
		return
	}
	fmt.Println(string(p))
	// 反序列化
	s := `{"name":"猪八戒","age":800}`
	var p1 person
	err = json.Unmarshal([]byte(s), &p1)
	if err != nil {
		fmt.Printf("marshal err: %v", err)
		return
	}
	fmt.Println(p1.Name)
}
