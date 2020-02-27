package main

import (
	"fmt"
	"strings"
)

type String string

type StrArray []string

func (s String) split(sep string) StrArray {
	return strings.Split(string(s), sep)
}

func (s StrArray) reverse() StrArray {
	runes := s
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return runes
}

func (s StrArray) join(sep string) String {
	return String(strings.Join(s, sep))
}
func main() {
	a := String("a:b:c:d")
	m := a.split(":").reverse().join("-")
	fmt.Println(m)
}
