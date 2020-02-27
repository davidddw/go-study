package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
