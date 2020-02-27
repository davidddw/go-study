package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type UnixProcess struct {
	pid   int
	ppid  int
	state rune
	pgrp  int
	sid   int

	binary string
}

func main() {
	dataBytes, err := ioutil.ReadFile("stat")
	if err != nil {
		return
	}
	p := &UnixProcess{pid: 1823}
	data := string(dataBytes)
	binStart := strings.IndexRune(data, '(') + 1
	binEnd := strings.IndexRune(data[binStart:], ')')
	p.binary = data[binStart : binStart+binEnd]
	data = data[binStart+binEnd+2:]
	fmt.Println(data)
	_, err = fmt.Sscanf(data,
		"%c %d %d %d",
		&p.state,
		&p.ppid,
		&p.pgrp,
		&p.sid)

	fmt.Println(p)
}
