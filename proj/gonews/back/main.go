package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/davidddw/go-common/logger"
	"github.com/davidddw/go-study/proj/gonews/back/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var client *redis.Client
var act string

var lock sync.Mutex

func main() {
	defer logger.Close()
	StartServ(":8081")
}

// StartServ start server at 8080
func StartServ(port string) {
	r := gin.Default()

	// r.LoadHTMLGlob("dist/*.html")        // 添加入口index.html
	// r.LoadHTMLFiles("dist/*/*")          // 添加资源路径
	// r.Static("/static", "./dist/static") // 添加资源路径
	// r.StaticFile("/", "dist/index.html") //前端接口

	r.Use(Cors())

	r.GET("/api/news", getNewsByPage)
	r.GET("/api/pull", pullNewsFromGithub)
	r.GET("/api/job/status", getJobStatus)
	r.Run(port)
}

func getNewsByPage(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageNum, _ := strconv.ParseInt(page, 10, 64)
	size := c.DefaultQuery("size", "10")
	pageSize, _ := strconv.ParseInt(size, 10, 64)
	news, count, err := common.GetPagedNews(pageNum, pageSize)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": count,
		"per":   pageSize,
		"items": news,
	})
}

func pullNewsFromGithub(c *gin.Context) {
	job := common.NewJob()
	_, err := job.CacheJob()
	go func() {
		lock.Lock()
		defer lock.Unlock()
		err := common.InitDataPuller()
		if err != nil {
			job.SetErr(err)
		} else {
			job.SetFinish()
		}
		job.CacheJob()
	}()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"jobid":  job.ID,
			"status": job.Status,
			"err":    job.Err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"jobid":  job.ID,
			"status": job.Status,
			"err":    nil,
		})
	}
}

func getJobStatus(c *gin.Context) {
	jobid := c.DefaultQuery("id", "1")
	job := &common.Job{ID: jobid}
	err := job.GetCacheJob()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jobid":  job.ID,
		"status": job.Status,
		"err":    job.Err.Error(),
	})
}

// Cors cors 中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
