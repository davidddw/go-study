package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/*")

	v1 := r.Group("/v1")
	{
		v1.POST("/upload", uploadFunc)
		v1.POST("/multi/upload", uploadMultiFunc)
	}

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	r.Run(":8080")
}

func uploadFunc(c *gin.Context) {
	file, _ := c.FormFile("upload")
	log.Println(file.Filename)

	dst := "static/uploadfile/" + file.Filename

	// 上传文件到指定的路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func uploadMultiFunc(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := "static/uploadfile/" + file.Filename
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, "Uploaded...")
}
