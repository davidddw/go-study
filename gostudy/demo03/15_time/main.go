package main

import (
	"fmt"
	"time"
)

// TIMELAYOUT timeformat
const TIMELAYOUT = "2006-01-02 15:04:05"

func main() {
	//取得东八区时间
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Minute())
	fmt.Println(now.Hour())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Printf("current timestamp1: %v\n", now.Unix())
	fmt.Printf("current timestamp1: %v\n", now.UnixNano())

	// time.Unix
	ret := time.Unix(1580554846, 0)
	fmt.Println(ret)
	fmt.Println(ret.String())
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	// 时间间隔
	fmt.Println(time.Second)

	// 时间+1小时
	fmt.Println(now.Add(time.Hour * 24))

	last, err := time.Parse("2006-01-02 15:04:05", "2020-02-05 12:00:00")
	if err != nil {
		fmt.Printf("parse time failed, err:%s.\n", err)
		return
	}
	fmt.Println(last)
	fmt.Println(now.Sub(last))

	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// 时间格式化
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))

	// 解析字符串时间
	timeObj, err := time.Parse("2006-01-02 15:04:05", "2020-02-01 19:00:00")
	if err != nil {
		fmt.Printf("parse time failed, err:%s.\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load location failed, err:%s.\n", err)
		return
	}
	localTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-02-01 19:00:00", loc)
	fmt.Println(localTime)
	fmt.Println(localTime.Unix())
}
