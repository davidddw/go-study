package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	userkey = "user"
)

func main() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	r.LoadHTMLGlob("templates/*")
	v1 := r.Group("/v1")
	{
		v1.GET("/book", getFunc)
		v1.POST("book", postFunc)
		v1.PUT("/book", putFunc)
		v1.DELETE("/book", deleteFunc)
		v1.PATCH("/book", patchFunc)
	}
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index",
		})
	})
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
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("age")
		nick := c.DefaultPostForm("name", "kim")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}

	r.Run("0.0.0.0:8080")
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
