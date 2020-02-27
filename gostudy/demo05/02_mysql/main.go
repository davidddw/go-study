package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const dsn = "cloud:passwd@tcp(192.168.1.190:3306)/mydb?charset=utf8"

var (
	db *sql.DB
)

type user struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func initDB() (err error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

func queryOne(id int) {
	var u1 user
	sqlStr := `select id, name, age from user where id=?`
	err := db.QueryRow(sqlStr, id).Scan(&u1.ID, &u1.Name, &u1.Age)
	ThrowError(err)
	fmt.Printf("u1:%#v\n", u1)
}

func queryMany() (users []*user, err error) {
	sqlStr := `select id, name, age from user `
	rows, err := db.Query(sqlStr)
	ThrowError(err)
	defer rows.Close()
	columns, err := rows.Columns()
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		ThrowError(err)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		var u user
		u.ID, _ = strconv.Atoi(record["id"])
		u.Name = record["name"]
		u.Age, _ = strconv.Atoi(record["age"])
		users = append(users, &u)
	}
	return
}

func queryManyAdv() (users []*user, err error) {
	sqlStr := `select id, name, age from user `
	rows, err := db.Query(sqlStr)
	ThrowError(err)
	defer rows.Close()
	columns, err := rows.Columns()
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Printf("scan failed, err：%v\n", err)
			return
		}
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		u := convertMapToObject(&record)
		users = append(users, &u)
	}
	return
}

func convertMapToObject(record *map[string]string) user {
	var u user
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(&u).Elem() // 为了改变对象的内部值，需使用引用
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		fieldName := t.Field(i).Tag.Get("db")
		if f.Kind() == reflect.Int {
			val, _ := strconv.Atoi((*record)[fieldName]) // 通过tag获取列数据
			f.SetInt(int64(val))
		} else if f.Kind() == reflect.String {
			f.SetString((*record)[fieldName])
		}
	}
	return u
}

func insert() int64 {
	sqlStr := `insert into user(name, age) values("lisi", 28)`
	res, err := db.Exec(sqlStr)
	ThrowError(err)

	id, err := res.LastInsertId()
	ThrowError(err)
	return id
}

func insertAdv(name string, age int) int64 {
	sqlStr := `insert into user(name, age) values(?, ?)`
	stmt, err := db.Prepare(sqlStr)
	ThrowError(err)
	defer stmt.Close()

	res, err := stmt.Exec(name, age)
	ThrowError(err)

	id, err := res.LastInsertId()
	ThrowError(err)
	return id
}

func update(id int, name string, age int) int64 {
	sqlStr := `update user set name=?, age=? where id=?`
	res, err := db.Exec(sqlStr, name, age, id)
	ThrowError(err)

	newID, err := res.RowsAffected()
	ThrowError(err)
	return newID
}

func updateAdv(id int, name string, age int) int64 {
	sqlStr := `update user set name=?, age=? where id=?`
	stmt, err := db.Prepare(sqlStr)
	ThrowError(err)
	defer stmt.Close()

	res, err := stmt.Exec(name, age, id)
	ThrowError(err)

	newID, err := res.RowsAffected()
	ThrowError(err)
	return newID
}

func deleteAdv(id int) int64 {
	sqlStr := `DELETE FROM user where id=?`
	stmt, err := db.Prepare(sqlStr)
	ThrowError(err)
	defer stmt.Close()

	res, err := stmt.Exec(id)
	ThrowError(err)

	newID, err := res.RowsAffected()
	ThrowError(err)
	return newID
}

// ThrowError ThrowError
func ThrowError(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func main() {
	err := initDB()
	ThrowError(err)
	fmt.Println("数据库连接成功！")
	//queryOne(2)
	//fmt.Println(insert())
	users, _ := queryManyAdv()
	for _, v := range users {
		fmt.Printf("user's info id:%v, name:%v, age:%v\n", v.ID, v.Name, v.Age)
	}
	//fmt.Println(deleteAdv(5))
	//fmt.Println(updateAdv(5, "zhangsan", 40))
	//fmt.Println(insertAdv("zhangsan", 40))
}
