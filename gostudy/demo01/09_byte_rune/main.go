package main

import (
	"fmt"
)

// fmt占位符
func main() {
	// s := "hello吉林智明創発"
	// fmt.Println(s)
	// n := len(s)
	// fmt.Println(n)

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%c\n", s[i])
	// }
	// fmt.Println("--------------------------")
	// for _, c := range s {
	// 	fmt.Printf("%c\n", c)
	// }
	// fmt.Println("===========================")
	// r := []rune(s)
	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("%c\n", r[i])
	// }

	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

}
