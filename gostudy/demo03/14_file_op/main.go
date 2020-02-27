package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	srcName = "t.txt"
	tmpName = "~tmp.txt"
)

func f() {
	in, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		os.Exit(-1)
	}
	defer in.Close()
	// 临时文件
	out, err := os.OpenFile(tmpName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", tmpName, err)
		os.Exit(-1)
	}
	defer out.Close()
	br := bufio.NewReader(in)
	index := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		newLine := strings.Replace(string(line), "b", "e", -1)
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
		fmt.Println("done ", index)
		index++
	}
	fmt.Println("FINISH!")
}

func main() {
	f()
	os.Rename(tmpName, srcName)
}
