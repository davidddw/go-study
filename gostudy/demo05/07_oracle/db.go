package main

import (
	"database/sql"
	"fmt"
	"log"

	maps "github.com/mitchellh/mapstructure"
)

// User model
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// InsertUser insert
func InsertUser(db *sql.DB, name string, age int) int64 {
	str := `insert into user1 (name, age) values(:1, :2)`
	res, err := sqlExec(db, str, name, age)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastInsertID
}

// DeleteUser delete
func DeleteUser(db *sql.DB, id int) int64 {
	str := `delete from user1 where id = :1`
	res, err := sqlExec(db, str, id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return rowsAffected
}

// UpdateUser update
func UpdateUser(db *sql.DB, name string, age int, id int) int64 {
	str := `update user1 set name=:1, age=:2 where id = :1`
	res, err := sqlExec(db, str, name, age, id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return rowsAffected
}

// SelectUserByCondition selectone
func SelectUserByCondition(db *sql.DB, id int) int {
	str := `select id, name, age from user1 where id = :1`
	users, err := sqlQuery(db, str, id)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Printf("%d\t%s\t%d\n", u.ID, u.Name, u.Age)
	}
	return len(users)
}

// SelectAllUser selectall
func SelectAllUser(db *sql.DB) int {
	str := `select id, name, age from user1`
	users, err := sqlQuery(db, str)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Printf("%d\t%s\t%d\n", u.ID, u.Name, u.Age)
	}
	return len(users)
}

// SelectCount selectall
func SelectCount(db *sql.DB) int64 {
	str := `select count(*) from user1`
	count, err := sqlQueryCount(db, str)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func sqlQuery(db *sql.DB, sqlStmt string, args ...interface{}) (users []*User, err error) {
	rows, err := db.Query(sqlStmt, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	// get columns
	columns, err := rows.Columns()
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(columns))

	// generate scanArgs
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return
		}
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		u := new(User)
		convertMapToUser(&record, u)
		users = append(users, u)
		//log.Printf("row[%d], name=[%s], age=[%d]\n", u.ID, u.Name, u.Age)
	}

	err = rows.Err()
	if err != nil {
		return
	}
	//log.Println("SQL Query success rows queried")
	return
}

func sqlQueryCount(db *sql.DB, sqlStmt string) (int64, error) {
	var count int64
	err := db.QueryRow(sqlStmt).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func sqlExec(db *sql.DB, sqlStmt string, params ...interface{}) (sql.Result, error) {
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(params...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}

func convertMapToUser(data *map[string]interface{}, out interface{}) {
	cfg := &maps.DecoderConfig{
		Metadata: nil,
		Result:   out,
		TagName:  "db",
	}
	decoder, _ := maps.NewDecoder(cfg)
	decoder.Decode(*data)
}
