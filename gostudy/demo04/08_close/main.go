package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 10
	ch1 <- 20
	ch1 <- 30
	close(ch1)
	for ret := range ch1 {
		fmt.Println(ret)
	}
}
