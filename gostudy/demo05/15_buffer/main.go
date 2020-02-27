package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var n int = 100000
	intToBytes := intToBytes(n)
	for _, v := range intToBytes {
		fmt.Printf("%d %#[1]v\n", v)
	}
	fmt.Println(intToBytes)
	fmt.Println(bytes.NewBufferString("10000"))
	testBufferString()
}

func intToBytes(n int) []byte {
	x := int32(n)
	//创建一个内容是[]byte的slice的缓冲器
	//与bytes.NewBufferString("")等效
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func testBufferString() {
	buf1 := bytes.NewBufferString("swift")
	buf2 := bytes.NewBuffer([]byte("swift"))
	buf3 := bytes.NewBuffer([]byte{'s', 'w', 'i', 'f', 't'})
	fmt.Println("===========以下buf1,buf2,buf3等效=========")
	fmt.Println("buf1:", buf1)
	fmt.Println("buf2:", buf2)
	fmt.Println("buf3:", buf3)
	fmt.Println("===========以下创建空的缓冲器等效=========")
	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println("buf4:", buf4)
	fmt.Println("buf5:", buf5)
}
