package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int

var wg sync.WaitGroup

func noBufChannel() {
	b = make(chan int) //不带缓冲区的通道
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从管道中取得：", x)
	}()
	b <- 10
	fmt.Println("10发送到goroutine从管道中")
	wg.Wait()
}

func bufChannel() {
	b = make(chan int, 1) //带缓冲区的通道
	b <- 10
	x := <-b
	fmt.Println("从goroutine管道中取得", x)
	close(b)
}

func main() {
	bufChannel()
}
