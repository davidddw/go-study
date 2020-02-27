package main

import (
	"fmt"
	"strings"
)

// fmt占位符
func main() {
	path := "C:\\Windows\\System32"
	fmt.Println(path)
	s2 := `
	asdfasdf
asdf
	dsaf`
	fmt.Println(s2)
	s3 := `C:\Windows\System32`
	fmt.Println(s3)

	//字符串操作
	fmt.Println(len(s3))

	// 拼接
	s3 += `\driver`
	fmt.Println(s3)
	s4 := fmt.Sprintf("%s\\%s", path, "driver")
	fmt.Println(s4)
	//分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(s4, "System64"))
	fmt.Println(strings.Contains(s4, "System32"))
	//前缀
	fmt.Println(strings.HasPrefix(s4, "D:"))
	fmt.Println(strings.HasPrefix(s4, "C:"))
	//后缀
	fmt.Println(strings.HasSuffix(s4, "driver1"))
	fmt.Println(strings.HasSuffix(s4, "driver"))
	s5 := "abcdefgcd"
	fmt.Println(strings.Index(s5, "c"))
	fmt.Println(strings.LastIndex(s5, "c"))
	//拼接
	fmt.Println(strings.Join(ret, "/"))
}
