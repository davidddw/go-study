package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/davidddw2017/panzer/gostudy/demo06/14_restdemo/asset"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	Id      int    `json:"id" gorm:"auto-increment"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
}

const (
	MyUser   = "cloud"
	Password = "passwd"
	Host     = "192.168.1.61"
	MPort    = 3306
	DbName   = "mydb"
	SUCCESS  = "success"
	Port     = "8080"
)

func mysqlConnectString(driver, host string) (db *gorm.DB, err error) {
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", MyUser, Password, host, MPort, DbName)
	db, err = gorm.Open(driver, connArgs)
	return
}

func sqliteConnectString(driver, host string) (db *gorm.DB, err error) {
	db, err = gorm.Open(driver, host)
	return
}

func dbConn(host string) (db *gorm.DB) {
	//db, err := mysqlConnectString("mysql", host)
	db, err := sqliteConnectString("sqlite3", "gorm.db")

	if err != nil {
		panic(err.Error())
	}
	db.Table("t_employee").AutoMigrate(&Employee{})
	return db
}

func GetEmployeeByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Query("id")
		var employee Employee
		if err := db.Table("t_employee").Where("id = ?", id).First(&employee).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Fprintln(gin.DefaultWriter, err)
		} else {
			c.JSON(200, employee)
		}
	}
}

func DeleteEmployeeByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		if err := db.Table("t_employee").Delete(&Employee{}, "id = ?", id).Error; err != nil {
			c.AbortWithStatus(500)
			fmt.Fprintln(gin.DefaultWriter, err)
		} else {
			c.JSON(200, SUCCESS)
		}
	}
}

func AddEmployee(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		address := c.PostForm("address")
		age, _ := strconv.Atoi(c.PostForm("age"))
		var employee = Employee{Name: name, Address: address, Age: age}
		if err := db.Table("t_employee").Create(&employee).Error; err != nil {
			c.AbortWithStatus(500)
			fmt.Fprintln(gin.DefaultWriter, err)
		} else {
			c.JSON(200, SUCCESS)
		}
	}
}

func UpdateEmployeeByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		name := c.PostForm("name")
		address := c.PostForm("address")
		age, _ := strconv.Atoi(c.PostForm("age"))
		var employee = Employee{id, name, address, age}
		if err := db.Table("t_employee").Save(&employee).Error; err != nil {
			c.AbortWithStatus(500)
			fmt.Fprintln(gin.DefaultWriter, err)
		} else {
			c.JSON(200, SUCCESS)
		}
	}
}

func GetAllEmployees(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var employees []Employee
		if err := db.Table("t_employee").Find(&employees).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Fprintln(gin.DefaultWriter, err)
		} else {
			c.JSON(200, employees)
		}
	}
}

func getTemplate() *template.Template {
	bytes, err := asset.Asset("view/employee.tmpl")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	t, _ := template.New("Index").Parse(string(bytes))
	return t
}

func setResource(router *gin.Engine) {
	router.SetHTMLTemplate(getTemplate())
	fs := assetfs.AssetFS{
		Asset:     asset.Asset,
		AssetDir:  asset.AssetDir,
		AssetInfo: asset.AssetInfo,
	}
	router.StaticFS("/vender", &fs)
}

func main() {
	var host string
	if len(os.Args) > 1 {
		host = os.Args[1]
	} else {
		host = Host
	}

	db := dbConn(host)
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	setResource(router)
	//router.LoadHTMLGlob("view/*")
	//router.Static("/vender/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Index", gin.H{
			"title": "Posts",
		})
	})
	v1 := router.Group("/rest/employee")
	{
		v1.GET("/listEmployee", GetAllEmployees(db))
		v1.GET("/getEmployee", GetEmployeeByID(db))
		v1.POST("/addEmployee", AddEmployee(db))
		v1.POST("/deleteEmployee", DeleteEmployeeByID(db))
		v1.POST("/updateEmployee", UpdateEmployeeByID(db))
	}
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	fmt.Println("Listening and serving HTTP on" + ":" + port)
	router.Run(":" + port)
}
