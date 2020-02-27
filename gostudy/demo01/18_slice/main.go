package main

import "fmt"

func main() {
	//定义
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//切片初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"吉林", "大连", "成都"}
	fmt.Println(s1, s2)
	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 2, 4}
	s3 := a1[0:4] //左包含右不包含
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4]
	s6 := a1[3:]
	s7 := a1[:]
	fmt.Println(s5, s6, s7)
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))

	// 再切片
	s8 := s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
	//切片是一个引用类型，都指向了底层的一个数组
	fmt.Println(s6)
	a1[6] = 2000 //修改底层数组的值
	fmt.Println(s6)
}
