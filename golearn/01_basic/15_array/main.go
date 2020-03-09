package main

import "fmt"

//数组
func main() {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("a1 type %T - value %v \na2 type %T - value %v\n", a1, a1, a2, a2)

	// 初始化方式1 默认零值
	arr1 := [3]bool{true, true, false}
	fmt.Println(arr1)
	// 初始化方式2
	// 自动推断数组长度
	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr2)
	// 初始化方式3
	// 根据索引方式来初始化
	arr3 := [...]int{0: 1, 4: 5}
	fmt.Println(arr3)

	city := [3]string{"北京", "上海", "广州"}
	// 根据索引遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	// for range遍历
	for i, v := range city {
		fmt.Println(i, v)
	}
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//多维数组遍历
	for _, v := range a11 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
