package common

import (
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/davidddw/go-common/logger"
	"github.com/davidddw/go-study/proj/gonews/back/config"
)

// News news structure
type News struct {
	ID    int64
	Title string
	Link  url.URL
	Ctime time.Time
}

func init() {
	err := logger.NewLogger("default")
	if err != nil {
		log.Fatal(err)
	}
}

func execute(workpath, path string, args ...string) ([]byte, error) {
	cmd := exec.Command(path, args...)
	cmd.Dir = workpath
	return cmd.Output()
}

// InitDataPuller load data from file and save into redis
func InitDataPuller() error {
	dir := config.SysConfig.Common.DataFolder
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	folderPath, _ := filepath.Abs(dir + "/news")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		out, err1 := execute(dir, "git", "clone", "https://github.com/gocn/news.git")
		if err1 != nil {
			logger.Errorf("%s", "clone failed")
			logger.Errorf("%s", err)
			logger.Errorf("%s", out)
			return err1
		} else {
			logger.Infof("%s", "Success to clone news")
		}
	}
	out, err := execute(dir, "git", "pull", "origin", "master")
	if err != nil {
		logger.Errorf("%s", "Pull failed")
		logger.Errorf("%s", err)
		logger.Errorf("%s", out)
		return err
	} else {
		logger.Infof("%s", "Success to pull news")
	}

	// 缓存数据操作
	files := GetFileList(dir)
	for _, file := range files {
		wg.Add(1)
		go CacheNews(file)
	}
	wg.Wait()
	logger.Infof("%s", "Success to cache news")

	return nil
}
