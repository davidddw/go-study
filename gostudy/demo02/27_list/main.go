package main

import (
	"container/list"
	"fmt"
)

func main() {
	lst := list.New()
	for i := 0; i < 5; i++ {
		lst.PushBack(i)
	}
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()
	fmt.Println(lst.Front().Value)
	fmt.Println(lst.Back().Value)
	lst.InsertAfter(6, lst.Front())
	fmt.Println()
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,061234
	}
	fmt.Println()
	fmt.Println()
	lst.MoveToFront(lst.Back()) //将尾部元素移动到首部
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,460123
	}
	fmt.Println()
	fmt.Println()

	l2 := list.New()
	l2.PushBackList(lst) //将l中元素放在l2的末尾
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出l2的值,460123
	}
	fmt.Println()
	lst.Init()           //清空l
	fmt.Print(lst.Len()) //0
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,无内容
	}

}
