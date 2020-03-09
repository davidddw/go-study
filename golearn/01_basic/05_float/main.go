package main

import (
	"fmt"
	"math"
)

func main() {
	var i1 = math.MaxFloat64
	fmt.Printf("%f\n", i1)
	i2 := math.MaxFloat32
	fmt.Printf("%f\n", i2)

	f1 := 1.23456
	fmt.Printf("%T\n", f1)
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
	f1 = float64(f2)
	fmt.Printf("%T\n", f1)

}
