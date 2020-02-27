package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// int to str
	i := int32(65536)
	fmt.Println(string(i))         //错误方式
	retStr := fmt.Sprintf("%d", i) //正确方式
	fmt.Printf("%#v %T\n", retStr, retStr)
	// int to str
	retStr1 := strconv.Itoa(int(i))
	fmt.Printf("%#v %T\n", retStr1, retStr1)

	// 从字符串中解析出对应的数据
	str := "100000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parse int failed, error: ", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)

	// string to int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	// 从字符串中解析出对应的bool
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 从字符串中解析出对应的float
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)
}
