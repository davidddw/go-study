package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("index.html")
	w.Write(buf)
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/posts/Go/15_socket", process)
	http.HandleFunc("/xxx/", f2)
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
