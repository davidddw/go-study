package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://192.168.1.200:9090/xxx")
	if err != nil {
		fmt.Printf("get url failed, err:%v.\n", err)
		return
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body failed, err:%v.\n", err)
		return
	}
	fmt.Println(string(ret))

}
