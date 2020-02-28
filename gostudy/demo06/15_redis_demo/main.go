package main

import (
	"fmt"
	"strconv"
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
	fmt.Println(generateKey(333, 6))
	fmt.Println(parseData("20120202033", 8))
}

func generateKey(data int64, length int) string {
	tmp := strconv.FormatInt(data, 10)
	prefixCount := length - len(tmp)
	target := make([]byte, length)
	for i := 0; i < length; i++ {
		if i < prefixCount {
			target[i] = '0'
		} else {
			target[i] = tmp[i-prefixCount]
		}
	}
	return string(target)
}

func parseData(data string, length int) (string, string) {
	subByte := []byte(data)
	if len(data) <= length {
		return "", ""
	}
	fmt.Println(len(data) - length)
	return string(subByte[:length]), string(subByte[length:])
}
