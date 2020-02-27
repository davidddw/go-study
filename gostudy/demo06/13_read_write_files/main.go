package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println(read0())
	fmt.Println(read1())
	fmt.Println(read2())
	fmt.Println(read3())
	write0()
	write1()
	write2()
	write3()
}

// ioutil.ReadFile
func read0() string {
	f, err := ioutil.ReadFile("stat")
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}

// 先从文件读取到file中，在从file读取到buf, buf在追加到最终的[]byte
func read1() string {
	f, err := os.Open("stat")
	if err != nil {
		fmt.Println("read fail")
		return ""
	}
	defer f.Close()
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		//从file读取到buf中
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return ""
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}

	return string(chunk)
}

// 先从文件读取到file, 在从file读取到Reader中，从Reader读取到buf, buf最终追加到[]byte
func read2() string {
	f, err := os.Open("stat")
	if err != nil {
		fmt.Println("read fail")
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var chunks []byte

	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

//读取到file中，再利用ioutil将file直接读取到[]byte中, 这是最优
func read3() string {
	f, err := os.Open("stat")
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

// 使用 io.WriteString 写入文件
func write0() error {
	var f *os.File
	fileName := "test1"
	strTest := "测试测试"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			fmt.Println("file create fail")
			return err
		}
	} else {
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		if err != nil {
			fmt.Println("file open fail", err)
			return err
		}
	}

	//将文件写进去
	n, err := io.WriteString(f, strTest)
	if err != nil {
		fmt.Println("write error", err)
		return err
	}
	fmt.Println("写入的字节数是：", n)
	return nil
}

// 使用 ioutil.WriteFile 写入文件
func write1() error {
	fileName := "test2"
	strTest := "测试测试"
	var d = []byte(strTest)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write fail")
		return err
	}
	fmt.Println("write success")
	return nil
}

// 使用 File(Write,WriteString) 写入文件
func write2() error {
	fileName := "test3"
	strTest := "测试测试"

	f, err := os.Create(fileName) //创建文件
	if err != nil {
		fmt.Println("create file fail")
		return err
	}
	defer f.Close()

	n2, err := f.Write([]byte(strTest)) //写入文件(字节数组)
	if err != nil {
		return err
	}
	fmt.Printf("写入 %d 个字节\n", n2)

	n3, err := f.WriteString("writesn") //写入文件(字节数组)
	if err != nil {
		return err
	}
	fmt.Printf("写入 %d 个字节\n", n3)

	f.Sync()
	return nil
}

// 使用 bufio.NewWriter 写入文件
func write3() error {
	fileName := "test4"
	f, err := os.Create(fileName) //创建文件
	if err != nil {
		fmt.Println("create file fail")
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err := w.WriteString("buffered\n")
	if err != nil {
		return err
	}
	fmt.Printf("写入 %d 个字节\n", n4)
	defer w.Flush()
	return nil
}
