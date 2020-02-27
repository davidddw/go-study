package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启3个goroutine
	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}
	//开启5个goroutine
	for j := 1; j <= 50; j++ {
		jobs <- j
	}
	close(jobs)
	//输出结果
	for a := 1; a <= 50; a++ {
		<-results
	}
}
