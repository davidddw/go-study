package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

// IniFile config file
const IniFile = "conf.ini"

func main() {
	var cfg Config
	err := LoadIni(IniFile, &cfg)
	dsn := fmt.Sprintf("%s/%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Sid)

	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		println("ok")
		log.Fatal(err)
	}

	//InsertUser(db, "智明", 15)
	//DeleteUser(db, 41)
	//UpdateUser(db, "智明", 30, 42)
	//fmt.Println(SelectCount(db))
	// SelectAllUser(db)
	SelectUserByCondition(db, 21)
	SelectAllUser(db)
}
