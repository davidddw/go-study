package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/davidddw/go-common/logger"
)

// News news structure
type News struct {
	ID    int64
	Title string
	Link  url.URL
	Ctime time.Time
}

var wg sync.WaitGroup

// InitDataPuller load data from file and save into redis
func InitDataPuller(dir string) {
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Dir = dir
	out, err := cmd.Output()
	logger.Infof("%s", string(out))

	if err != nil {
		logger.Errorf("%s", "Pull failed")
		logger.Errorf("%s", err)
		logger.Errorf("%s", out)
	} else {
		if !strings.Contains(string(out), "Already up to date") {
			logger.Infof("%s", "Pull success")

			// 缓存数据操作
			files := getFileList(dir)
			for _, file := range files {
				wg.Add(1)
				go cacheNews(file)
			}
			wg.Wait()
			logger.Infof("%s", "Success to cache news")
		}
	}
}

func cacheNews(path string) {
	defer wg.Done()
	newsList, _ := parseNews(path)
	for _, item := range newsList {
		cache := map[string]interface{}{
			"id":    item.ID,
			"title": item.Title,
			"link":  item.Link.String(),
			"ctime": item.Ctime.Format("20060102"),
		}
		setNewsToCache(cache) // 缓存数据
	}
}

func setNewsToCache(cache map[string]interface{}) error {
	rootKey := CachePrefix
	sortedKey := SortedPrefix
	var key1, key2 string
	if value, ok := cache["ctime"].(string); ok {
		err := client.SAdd(rootKey, value).Err()
		if err != nil {
			return err
		}
		key1 = value
	}
	if value, ok := cache["id"].(int64); ok {
		key2 = fmt.Sprintf("%03d", value)
		err := client.SAdd(sortedKey, key1+key2).Err()
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

func getFileList(dir string) []string {
	files := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if strings.Contains(path, ".git") {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		baseName := filepath.Base(path)
		if !strings.Contains(baseName, "-") {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".md" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files
}

func parseNews(path string) (newsList []News, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`(?i)#{0,3}\s*GoCN每日新闻\(([\d-]*)\)\n+\D*((.*\n)+?\n)`)
	allMatch := reg.FindAllSubmatch(data, -1)
	for _, item := range allMatch {
		loc, _ := time.LoadLocation("Local")
		ctime, err := time.ParseInLocation("2006-01-02", string(item[1]), loc)
		if err != nil {
			continue
		}
		subReg := regexp.MustCompile(`(\d)\.\s*(.*)\s*(http.*)\n?`)
		subAll := subReg.FindAllSubmatch(item[2], -1)
		for _, subItem := range subAll {
			id, err := strconv.ParseInt(string(subItem[1]), 10, 64)
			if err != nil {
				continue
			}
			title := strings.TrimSpace(string(subItem[2]))
			sURL, err := url.Parse(strings.TrimSpace(string(subItem[3])))
			if err != nil {
				continue
			}
			singleNews := News{id, title, *sURL, ctime}
			newsList = append(newsList, singleNews)
		}
	}
	return
}
