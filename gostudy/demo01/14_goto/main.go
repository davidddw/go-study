package main

import "fmt"

// goto
func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; i < 'Z'; j++ {
			if j == 'C' {
				break
			}
			fmt.Printf("%v - %c\n", i, j)
		}
	}
	fmt.Println("+++++++++++++++++++++++++")
	for i := 0; i < 10; i++ {
		for j := 'A'; i < 'Z'; j++ {
			if j == 'C' {
				goto breakTag
			}
			fmt.Printf("%v - %c\n", i, j)
		}
	}
breakTag:
	fmt.Println("over")
}
