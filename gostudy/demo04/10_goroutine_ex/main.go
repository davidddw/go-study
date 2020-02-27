package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用goroutine和channel实现计算int64随机数和的程序
// 1. 开启goroutine循环生成随机数发送到jobChan
// 2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，结果发送resultChan
// 3. 从resultChan取出并打印

// Job 随机数
type Job struct {
	value int64
}

// Result 计算各位数的和
type Result struct {
	job *Job
	sum int64
}

var wg sync.WaitGroup
var once sync.Once

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generate(data chan<- *Job) {
	defer wg.Done()
	defer close(data) //避免deadlock
	var i = 0
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		data <- newJob
		time.Sleep(time.Millisecond * 500)
		i++
		// if i > 5 { //避免deadlock
		// 	break
		// }
	}

}

func calc(data <-chan *Job, resultChan chan<- *Result) {
	defer wg.Done()
	defer once.Do(func() { //避免deadlock
		close(resultChan)
	})
	for {
		job, ok := <-data
		if !ok { //避免deadlock
			break
		}
		n := job.value
		var sum int64
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		newResult := &Result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go generate(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go calc(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value: %d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}
