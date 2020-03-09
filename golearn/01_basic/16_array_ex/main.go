package main

import "fmt"

// 数组练习题
// 求数组的和
func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)

	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d %d)\n", i, j)
			}
		}
	}
}
