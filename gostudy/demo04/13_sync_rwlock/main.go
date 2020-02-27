package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var rwLock sync.RWMutex

func read() {
	rwLock.RLock()
	// lock.Lock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	// lock.Unlock()
	fmt.Println(x)
	wg.Done()
}

func write() {
	rwLock.Lock()
	// lock.Lock()
	x++
	time.Sleep(5 * time.Millisecond)
	rwLock.Unlock()
	// lock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
