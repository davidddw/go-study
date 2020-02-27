package main

import "fmt"

// panic

func funcA() {
	fmt.Println("A")
}

// recover()必须搭配defer使用
// defer一定要在可能引发panic之前定义
func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("error !")
	fmt.Println("B")
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
