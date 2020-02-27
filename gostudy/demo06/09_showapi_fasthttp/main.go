package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/davidddw/go-utils/showSdk"
)

const (
	timeFormat      = "2006-01-02 15:04:05.000"
	timeFormatShort = "2006-01-02 15:04:05"
	timeFormatFile  = "20060102"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormatShort)
}

// Result result value
type ShowAPIResult struct {
	ShowapiResError string  `json:"showapi_res_error"`
	ShowapiResID    string  `json:"showapi_res_id"`
	ShowapiResCode  int     `json:"showapi_res_code"`
	ShowapiResBody  ResBody `json:"showapi_res_body"`
}

// ResBody return value
type ResBody struct {
	RetCode     int    `json:"ret_code"`
	Contentlist []Unit `json:"list"`
}

// Unit Unit
type Unit struct {
	Prov string `json:"prov"`
	P90  string `json:"p90"`
	P0   string `json:"p0"`
	P95  string `json:"p95"`
	P97  string `json:"p97"`
	P98  string `json:"p98"`
	P89  string `json:"p89"`
	P92  string `json:"p92"`
	P93  string `json:"p93"`
	Ct   Time   `json:"ct"`
}

const (
	showapiUrl = "http://route.showapi.com/138-46"
	appid      = 146977
	sign       = "5bd7e8474b2643978656b1872f637670"
)

func main() {
	result := getOilPricesToday(showapiUrl)
	body := result.ShowapiResBody.Contentlist
	if len(body) != 0 {
		content := result.ShowapiResBody.Contentlist
		for _, v := range content {
			fmt.Printf("省份：%s\t92号：%s\t95号：%s\n", v.Prov, v.P92, v.P95)
		}
	} else {
		fmt.Println("没有了")
	}
	filename := fmt.Sprintf("今日油价(%s).xlsx", time.Now().Format(timeFormatFile))
	saveToExcel(filename, &result.ShowapiResBody.Contentlist)
}

func getOilPricesToday(url string) ShowAPIResult {
	currentTime := time.Now()
	res := showSdk.ShowAPIRequest(showapiUrl, appid, sign)
	ret, _ := res.Post()
	endTime := time.Now()
	during := endTime.Sub(currentTime)
	fmt.Printf("查询执行消耗的时间为:%v秒\n", during)
	var result ShowAPIResult
	err := json.Unmarshal([]byte(ret), &result)
	if err != nil {
		fmt.Println("JSON Parse ERR:", err)
	}
	return result
}

func saveToExcel(filename string, data *[]Unit) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	style, _ := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}]}`)
	f.SetCellStyle("Sheet1", "A1", "G32", style)
	f.SetColWidth("Sheet1", "A", "F", 12)
	f.SetColWidth("Sheet1", "G", "G", 20)
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "今日油价")
	f.SetCellValue("Sheet1", "B1", "89号汽油")
	f.SetCellValue("Sheet1", "C1", "92号汽油")
	f.SetCellValue("Sheet1", "D1", "95号汽油")
	f.SetCellValue("Sheet1", "E1", "98号汽油")
	f.SetCellValue("Sheet1", "F1", "0号柴油")
	f.SetCellValue("Sheet1", "G1", "更新日期")
	for i, v := range *data {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), v.Prov)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), v.P89)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), v.P92)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), v.P95)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), v.P98)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), v.P0)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), v.Ct.String())
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs(filename); err != nil {
		println(err.Error())
	}
}
