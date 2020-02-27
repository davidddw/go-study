package main

import (
	"fmt"

	"github.com/Lofanmi/pinyin-golang/pinyin"
)

func main() {
	a := "牛郎"
	b := "织女"
	fmt.Println(a, b)
	m1 := convHanToNumber(a + b)
	fmt.Println(m1)
	fmt.Println(calc(m1))
}

func convHanToNumber(han string) []byte {
	dict := pinyin.NewDict()
	strPinyin := dict.Convert(han, "").None()
	bytePinyin := []byte(strPinyin)
	for i := 0; i < len(bytePinyin); i++ {
		bytePinyin[i] -= 96
	}
	var str string
	for i := 0; i < len(bytePinyin); i++ {
		str += fmt.Sprintf("%d", bytePinyin[i])
	}
	b := []byte(str)
	for i := 0; i < len(b); i++ {
		b[i] -= 48
	}
	return b
}

func calc(num []byte) string {
	var tmpSlice []byte
	var tens, units byte
	tmpSlice = append(tmpSlice, num...)
	for i := 0; i < len(num)-2; i++ {
		var resultSlice = make([]byte, len(num))
		for j := 0; j < len(num)-i-1; j++ {
			resultSlice[j] = (tmpSlice[j] + tmpSlice[j+1]) % 10
		}
		fmt.Println(resultSlice)
		tmpSlice = resultSlice
		tens, units = resultSlice[0], resultSlice[1]
	}
	return fmt.Sprintf("%d%d%%\n", tens, units)
}
