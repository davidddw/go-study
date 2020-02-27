package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/davidddw2017/panzer/gostudy/demo06/07_showapi/showSdk/normalRequest"
)

// Result result value
type Result struct {
	ShowapiResError string  `json:"showapi_res_error"`
	ShowapiResID    string  `json:"showapi_res_id"`
	ShowapiResCode  int     `json:"showapi_res_code"`
	ShowapiResBody  ResBody `json:"showapi_res_body"`
}

// ResBody return value
type ResBody struct {
	AllPages    int    `json:"allPages"`
	RetCode     string `json:"ret_code"`
	Contentlist []Unit `json:"contentlist"`
}

// Unit Unit
type Unit struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	currentTime := time.Now()
	res := normalRequest.ShowapiRequest("http://route.showapi.com/1635-1", 146977, "5bd7e8474b2643978656b1872f637670")
	res.AddTextPara("page", "141")
	ret, _ := res.Post()
	var result Result
	err := json.Unmarshal([]byte(ret), &result)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}
	count := result.ShowapiResBody.AllPages
	if count > 0 {
		content := result.ShowapiResBody.Contentlist
		for _, v := range content {
			fmt.Printf("问题：%s\t\t\t答案：%s\n", v.Question, v.Answer)
		}
	} else {
		fmt.Println("没有了")
	}
	endTime := time.Now()
	during := endTime.Sub(currentTime)
	fmt.Println(during)
}
