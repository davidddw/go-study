package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func f1(i int) {
	defer wg.Done()
	ti := time.Millisecond * time.Duration(rand.Intn(300))
	time.Sleep(ti)
	fmt.Println("hello", i, "cost:", ti)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait() //等待计数器减为0
}
