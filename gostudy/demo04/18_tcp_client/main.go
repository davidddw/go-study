package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial 127.0.0.1:20000 failed, err:%v.\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	var msg string
	reader := bufio.NewReader(os.Stdin)
	// if len(os.Args) > 0 {
	// 	msg = os.Args[1]
	// } else {
	// 	msg = "hello"
	// }
	//conn.Write([]byte(msg))
	for {
		//fmt.Scanln(&msg)
		fmt.Print("请输入信息：")
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}
