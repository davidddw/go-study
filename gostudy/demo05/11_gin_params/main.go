package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	r.GET("/user", func(c *gin.Context) {
		firstname := c.DefaultQuery("name", "kim") // 获取query中的name，没有的话就为kim
		lastname := c.Query("age")                 // 获取query中的age

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	r.POST("/user/form_post", func(c *gin.Context) {
		message := c.PostForm("age")
		nick := c.DefaultPostForm("name", "kim")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.Run(":8080")
}
