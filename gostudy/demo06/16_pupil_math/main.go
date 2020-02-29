package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// Row 表格的行数
var Row = flag.Int("r", 26, "Input Rows Number")

// Col 表格的列数
var Col = flag.Int("c", 4, "Input Column Number")

// Col 表格的列数
var filename = flag.String("f", "001.md", "Input Column Number")

var answer100 [1024]int

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	generateQuestions(*filename, *Row, *Col)
}

func generateQuestions(filename string, row, col int) {
	output := generateTitle()
	output += generateLine(0, col)
	output += titleLine(col)
	for i := 1; i < row; i++ {
		output += generateLine(i, col)
	}
	ioutil.WriteFile(filename, []byte(output), 0644)
}

func generateTitle() string {
	ret := `<h2 style="text-align:center">小学生口算题 </h2>` + "\n"
	ret += `<h4 style="text-align:center">姓名：<u>&nbsp;
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</u> 日期：<u>&nbsp;&nbsp;
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</u> 月 <u>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
&nbsp;&nbsp;&nbsp;</u> 日   时间：<u>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</u>  对题：<u>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</u>道 </h4>` + "\n\n"
	return ret

}

func titleLine(col int) string {
	output := "|"
	for i := 0; i < col; i++ {
		output += " ---- |"
	}
	output += "\n"
	return output
}

func generateLine(line int, col int) string {
	var count = 0
	var output = "|"
	for count < col {
		a := rand.Intn(99) + 1
		b := rand.Intn(99) + 1
		op := rand.Intn(2)
		//fmt.Println(op)
		if op == 0 && (a+b) < 100 {
			answer100[line*col+count] = a + b
			output += fmt.Sprintf("%2d + %2d = | ", a, b)
			count++
		}
		if op == 1 && a > b && (a-b) > 0 {
			answer100[line*col+count] = a - b
			output += fmt.Sprintf("%2d - %2d = | ", a, b)
			count++
		}
	}
	output += "\n"
	return output
}
