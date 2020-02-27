package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read from conn failed, err:%v.\n", err)
			os.Exit(1)
		}
		fmt.Printf("收到客户端发来的数据：%s\n", string(buf[:n]))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err:%v.\n", err)
		return
	}
	defer listener.Close()
	for {
		fmt.Println("等待客户端来链接....")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v.\n", err)
			return
		}
		go process(conn)
	}
}
