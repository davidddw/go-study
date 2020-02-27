package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//writeDemo1()
	//writeDemo2()
	writeDemo3()
}

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	//fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	fileObj.Write([]byte("どうぞ宜(よろ)しくお願(ねが)いいたします\n"))
	fileObj.WriteString("こんにちは\n")
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	//fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("どうぞ宜(よろ)しくお願(ねが)いいたします\n")
	wr.WriteString("こんにちは\n")
	wr.Flush()
}

func writeDemo3() {
	str := "どうぞ宜(よろ)しくお願(ねが)いいたします\n"
	str += "こんにちは\n"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
}
