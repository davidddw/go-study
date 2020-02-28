package common

import (
	"strconv"
	"sync"
	"time"

	"github.com/davidddw/go-common/logger"
	"github.com/davidddw/go-study/proj/gonews/back/config"
	"github.com/go-redis/redis"
)

var (
	wg           sync.WaitGroup
	client       *redis.Client
	rootKey      string
	sortedPrefix string
)

func init() {
	if client == nil {
		client = newClient(config.SysConfig.Redis.Host, config.SysConfig.Redis.DB)
	}
	rootKey = config.SysConfig.Redis.CachePrefix
	sortedPrefix = config.SysConfig.Redis.SortedPrefix
}

func newClient(host string, db int) *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       db,
	})
	return client
}

func CacheNews(path string) {
	defer wg.Done()
	newsList, _ := ParseNews(path)
	for _, item := range newsList {
		cache := map[string]interface{}{
			"id":    item.ID,
			"title": item.Title,
			"link":  item.Link.String(),
			"ctime": item.Ctime.Format("20060102"),
		}
		SetNewsToCache(cache) // 缓存数据
	}
}

// SetNewsToCache store data to redis
func SetNewsToCache(cache map[string]interface{}) error {
	var key1, key2 string
	if value, ok := cache["ctime"].(string); ok {
		err := client.SAdd(rootKey, value).Err()
		if err != nil {
			return err
		}
		key1 = value
	}
	if value, ok := cache["id"].(int64); ok {
		key2 = GenerateKey(value, 8)
		err := client.SAdd(sortedPrefix, key1+key2).Err()
		if err != nil {
			return err
		}
		key1 := rootKey + ":" + key1
		err = client.SAdd(key1, key2).Err()
		if err != nil {
			return err
		}
		key2 = key1 + ":" + key2
		return client.HMSet(key2, cache).Err()
	}
	return nil
}

// GenerateKey return string which has 0 prefix
// for number 25 and length equals 6 => 000025
func GenerateKey(data int64, length int) string {
	tmp := strconv.FormatInt(data, 10)
	prefixCount := length - len(tmp)
	target := make([]byte, length)
	for i := 0; i < length; i++ {
		if i < prefixCount {
			target[i] = '0'
		} else {
			target[i] = tmp[i-prefixCount]
		}
	}
	return string(target)
}

// ParseData return two strings
// for string 202002020004 and length equals 8 => 20200202, 0004
func ParseData(data string, length int) (string, string) {
	subByte := []byte(data)
	if len(data) <= length {
		return "", ""
	}
	return string(subByte[:length]), string(subByte[length:])
}

// 获取新闻
func GetPagedNews(pageNum int64, pageSize int64) ([]map[string]string, int64, error) {
	start := time.Now()
	offset := (pageNum - 1) * pageSize
	sortedKey, err := client.Sort(sortedPrefix, &redis.Sort{Offset: offset, Count: pageSize, Order: "desc"}).Result()
	if err != nil {
		return nil, 0, err
	}
	count, err := client.SCard(sortedPrefix).Result()
	if err != nil {
		return nil, 0, err
	}
	var newsList []map[string]string
	for i := 0; i < len(sortedKey); i++ {
		length := len(sortedKey[i])
		if length != 0 {
			key1, id := ParseData(sortedKey[i], 8)
			news, err := GetNewsCache(rootKey + ":" + key1 + ":" + id)
			if err != nil {
				continue
			}
			newsList = append(newsList, news)
		}
	}
	end := time.Now()
	logger.Infof("time cost %v\n", end.Sub(start))
	return newsList, count, nil
}

// GetNewsCache 获取新闻缓存
func GetNewsCache(key string) (map[string]string, error) {
	return client.HGetAll(key).Result()
}

func CacheJob(cache map[string]interface{}) error {
	if value, ok := cache["id"].(string); ok {
		err := client.SAdd("jobs", value).Err()
		if err != nil {
			return err
		}
		key := "jobs-" + value
		return client.HMSet(key, cache).Err()
	}
	return nil
}

func GetJob(id string) (map[string]string, error) {
	key := "jobs-" + id
	return client.HGetAll(key).Result()
}
