package main

import "fmt"

// fmt

func main() {
	fmt.Print("北京")
	fmt.Print("吉林")
	fmt.Println()
	fmt.Println("北京")
	fmt.Println("吉林")

	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)

	m1 := make(map[string]int, 1)
	m1["吉林"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)

	fmt.Printf("%d%%\n", 100)
	fmt.Printf("%q\n", 65)

	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	s := "吉林智明"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	a := 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)

	n := 12.34
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.2f\n", n)
	fmt.Printf("%9.f\n", n)

	fmt.Printf("%9s\n", s)
	fmt.Printf("%-9s\n", s)
	fmt.Printf("%9.7s\n", s)
	fmt.Printf("%-9.7s\n", s)
	fmt.Printf("%9.2s\n", s)
	fmt.Printf("%09s\n", s)
}
