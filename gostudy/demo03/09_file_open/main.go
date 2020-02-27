package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	readFromFileByIoutil()
}

func readFromFileByIoutil() {
	// read file
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Print(string(ret))
}

func readFromFileByBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	// read file
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	// read file
	var buf = make([]byte, 128)
	for {
		n, err := fileObj.Read(buf)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		//fmt.Printf("read %d byte\n", n)
		fmt.Print(string(buf[:n]))
	}
}
