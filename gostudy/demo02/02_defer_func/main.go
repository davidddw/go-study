package main

import "fmt"

func demo() {
	fmt.Println("start")
	fmt.Println("hello world")
	fmt.Println("stop")
}

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("hello world")
	fmt.Println("stop")
}

func deferDemo2() {
	fmt.Println("start")
	defer fmt.Println("hello 1")
	defer fmt.Println("hello 2")
	defer fmt.Println("hello 3")
	fmt.Println("stop")
}

func main() {
	demo()
	fmt.Println("==============================")
	deferDemo()
	deferDemo2()
}
