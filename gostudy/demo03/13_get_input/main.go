package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	var s string
	fmt.Print("input string:")
	fmt.Scanln(&s)
	fmt.Printf("you input is:%s\n", s)
}

func userBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input string:")
	s, _ = reader.ReadString('\n')
	fmt.Printf("you input is:%s", s)
}

func main() {
	userBufio()
}
