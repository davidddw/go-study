package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/davidddw/go-common/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	// Dir file path
	Dir = "E:\\news"
	// CachePrefix file path
	CachePrefix = "gonews"
	// SortedPrefix sort path
	SortedPrefix = "gonews_sort"
	// Host redis host
	Host = "192.168.1.61:6379"
	// CacheDB redis db
	CacheDB = 2
)

var client *redis.Client
var act string

func init() {
	flag.StringVar(&act, "a", "cache", "the action to run service, values 'api' or 'cache'")
	if client == nil {
		client = newClient()
	}
	err := logger.NewLogger("default")
	if err != nil {
		log.Fatal(err)
	}
}

func newClient() *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     Host,
		Password: "",
		DB:       CacheDB,
	})
	return client
}

func main() {
	flag.Parse()
	if act == "api" {
		StartServ()
	} else {
		InitDataPuller(Dir)
	}
	defer logger.Close()
}

// StartServ start server at 8080
func StartServ() {
	r := gin.Default()

	// r.LoadHTMLGlob("dist/*.html")        // 添加入口index.html
	// r.LoadHTMLFiles("dist/*/*")          // 添加资源路径
	// r.Static("/static", "./dist/static") // 添加资源路径
	// r.StaticFile("/", "dist/index.html") //前端接口

	r.Use(Cors())

	r.GET("/api/news", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageNum, _ := strconv.ParseInt(page, 10, 64)
		size := c.DefaultQuery("size", "10")
		pageSize, _ := strconv.ParseInt(size, 10, 64)
		news, count, err := getPagedNews(pageNum, pageSize)

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
	})
	r.Run(":8081")
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

func getData(pageNum, pageSize int64) []byte {
	news, count, _ := getPagedNews(pageNum, pageSize)
	data := map[string]interface{}{
		"total":    count,
		"pagesize": pageSize,
		"items":    news,
	}
	jData, _ := json.Marshal(data)
	return jData
}

// 获取新闻
func getPagedNews(pageNum int64, pageSize int64) ([]map[string]string, int64, error) {
	start := time.Now()
	offset := (pageNum - 1) * pageSize
	sortedKey, err := client.Sort(SortedPrefix, &redis.Sort{Offset: offset, Count: pageSize, Order: "desc"}).Result()
	if err != nil {
		return nil, 0, err
	}
	count, err := client.SCard(SortedPrefix).Result()
	if err != nil {
		return nil, 0, err
	}
	var newsList []map[string]string
	for i := 0; i < len(sortedKey); i++ {
		length := len(sortedKey[i])
		if length != 0 {
			sByte := []byte(sortedKey[i])
			key1 := string(sByte[:8])
			id := string(sByte[length-3:])
			news, err := getNewsCache(CachePrefix + ":" + key1 + ":" + id)
			if err != nil {
				continue
			}
			newsList = append(newsList, news)
		}
	}
	end := time.Now()
	fmt.Printf("cost %v\n", end.Sub(start))
	return newsList, count, nil
}

// 获取全部新闻
func getAllNews(pageNum int64, pageSize int64) ([]map[string]string, int64) {
	start := time.Now()
	key0 := CachePrefix
	keys1, _ := client.SMembers(key0).Result()
	newsList := map[string]map[string]string{}
	var count int
	for _, key1 := range keys1 {
		keys2, _ := client.SMembers(key0 + ":" + key1).Result()
		for _, key2 := range keys2 {
			news, err := getNewsCache(key0 + ":" + key1 + ":" + key2)
			if err == nil {
				newsList[news["ctime"]+news["id"]] = news
			}
			count++
		}
	}
	fmt.Println(count)
	allNews := sortNews(newsList)
	pageNews := []map[string]string{}
	var i int64 = 0
	for _, item := range allNews {
		if i >= (pageNum-int64(1))*pageSize && i < pageNum*pageSize {
			pageNews = append(pageNews, item)
		}
		i++
	}
	end := time.Now()
	fmt.Printf("cost %v\n", end.Sub(start))
	return pageNews, i
}

// 对新闻进行排序
func sortNews(raw map[string]map[string]string) []map[string]string {
	keys := []string{}
	data := []map[string]string{}
	for key := range raw {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for _, key := range keys {
		data = append(data, raw[key])
	}
	return data
}

// 获取新闻缓存
func getNewsCache(key string) (map[string]string, error) {
	return client.HGetAll(key).Result()
}
