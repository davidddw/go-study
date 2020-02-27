package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "cloud:passwd@tcp(192.168.1.190:3306)/mydb?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open %s failed, err：%v\n", dsn, err)
		return
	}
	//关闭数据库
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err：%v\n", dsn, err)
		return
	}
	fmt.Println("数据库连接成功！")
}
