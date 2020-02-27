package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	m := 1
	fmt.Println(&m)
	for x := range ch1 {
		ch2 <- x * x
	}
	// for {
	// 	x, ok := <-ch1
	// 	if !ok {
	// 		break
	// 	}
	// 	ch2 <- x * x
	// }
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	f1(a)
	f2(a, b)
	f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
