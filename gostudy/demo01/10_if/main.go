package main

import (
	"fmt"
)

// if 判断
func main() {
	age1 := 19
	if age1 > 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年")
	}

	age2 := 19
	if age2 > 35 {
		fmt.Println("中年人")
	} else if age2 > 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年")
	}

	if age3 := 19; age3 > 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年")
	}

}
