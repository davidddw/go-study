package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Printf("a:%v len(a):%d cap(a):%d\n", a, len(a), cap(a))

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])
	fmt.Println(a1)
}
