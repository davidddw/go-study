package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadFunc(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "static/uploadfile/" + file.Filename

	// 上传文件到指定的路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func uploadMultipartFunc(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files[]"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := "static/uploadfile/" + file.Filename
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, "Uploaded...")
}

func getFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET",
	})
}

func postFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "POST",
	})
}

func patchFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PATCH",
	})
}

func putFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PUT",
	})
}

func deleteFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE",
	})
}
