package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-oci8"
)

const dsn = "oracle/ora123@192.168.1.190:1521/orcl"

type user struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	db, err := sql.Open("oci8", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := sqlExec(db, `insert into user1 (id, name, age) values(:1, :2, :3)`, 7, "Jerry", 20); err != nil {
		log.Fatal(err)
	}

	if err := sqlQuery(db, `select name, age from user1`); err != nil {
		log.Fatal(err)
	}
}

func sqlQuery(db *sql.DB, sqlStmt string) error {
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	var n int
	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			return err
		}
		n++
		log.Printf("row[%d], name=[%s], age=[%d]\n", n, name, age)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	log.Printf("SQL Query success rows queried %d\n", n)
	return nil
}

func sqlExec(db *sql.DB, sqlStmt string, params ...interface{}) error {
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(params...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	if lastInsertID > 0 {
		log.Printf("SQL Execute success rows affected %d\n", lastInsertID)
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num > 0 {
		log.Printf("SQL Execute success rows affected %d\n", num)
	}
	return nil
}
