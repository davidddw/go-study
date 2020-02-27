package dao

import (
	"fmt"
	"time"

	"github.com/davidddw2017/panzer/gostudy/demo06/06_gorm/table"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// MyUser database user name
	MyUser = "cloud"
	// Password database user password
	Password = "passwd"
	// Host database host ip
	Host = "192.168.1.61"
	// MPort database host port
	MPort = 3306
	// DbName database db name
	DbName = "mydb"
	// SUCCESS return value
	SUCCESS = "success"
	// Port listen port
	Port = "8080"
)

var _DB *gorm.DB

func DB() *gorm.DB {
	return _DB
}

func init() {
	_DB = initDB()
}

func initDB() *gorm.DB {
	// In our docker dev environment use
	// db, err := gorm.Open("mysql", "go_web:go_web@tcp(database:3306)/go_web?charset=utf8&parseTime=True&loc=Local")
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, MPort, DbName)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}

func CreatePerson(person *table.Person) (err error) {
	err = DB().Create(person).Error

	return
}

func GetPersonById(personId int64) (person *table.Person, err error) {
	person = new(table.Person)
	err = DB().Where("id = ?", personId).First(person).Error

	return
}

func GetAllPerson() (persons []*table.Person, err error) {
	err = DB().Find(&persons).Error
	return
}

func UpdatePersonNameById(name string, personId int64) (err error) {
	person := new(table.Person)
	err = DB().Where("id = ?", personId).First(person).Error
	if err != nil {
		return
	}

	person.Name = name
	err = DB().Save(person).Error

	return
}

func DeletePersonById(personId int64) (err error) {
	person := new(table.Person)
	err = DB().Where("id = ?", personId).First(person).Error
	if err != nil {
		return
	}
	err = DB().Delete(person).Error

	return
}
