package main

import (
	"encoding/json"
	"fmt"

	"github.com/davidddw2017/panzer/gostudy/demo06/08_showapi/showSdk"
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
	ShowapiCode int    `json:"showapi_fee_code"`
	AllNum      int    `json:"allNum"`
	MaxResult   int    `json:"maxResult"`
	CurrentPage int    `json:"currentPage"`
	Remark      string `json:"remark"`
}

// Unit Unit
type Unit struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	sa := showSdk.NewShowApi(146977, "5bd7e8474b2643978656b1872f637670")
	sa.Add("page", "1")
	ret, err := sa.Post("http://route.showapi.com/1635-1")
	if err != nil {
		fmt.Println("POST ERR:", err)
	}
	var result Result
	err = json.Unmarshal([]byte(ret), &result)
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
}
