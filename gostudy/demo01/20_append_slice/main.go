package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "广州"}
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	// s1[3] = "深圳" //索引越界
	// fmt.Println(s1)
	s1 = append(s1, "深圳") // append追加， 底层数组放不下， 底层数组换一个
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "吉林", "成都")
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))

	ss := []string{"武汉", "西安", "大连"}
	s1 = append(s1, ss...)
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
}
