package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

var mm = sync.Map{}

var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func load1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=%v, v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func load2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			lock.Lock()
			set(key, n)
			lock.Unlock()
			fmt.Printf("k=%v, v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func load3() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			mm.Store(key, n)
			value, _ := mm.Load(key)
			fmt.Printf("k=%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	// load1()
	//load2()
	load3()
}
