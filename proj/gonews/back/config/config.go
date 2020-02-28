package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type Config struct {
	Common `ini:"common"`
	Redis  `ini:"redis"`
}

type Common struct {
	DataFolder string `ini:"dataFolder"`
	Port       int    `ini:"port"`
	Repo       string `ini:"repo"`
}

type Redis struct {
	Host         string `ini:"host"`
	DB           int    `ini:"db"`
	CachePrefix  string `ini:"cachePrefix"`
	SortedPrefix string `ini:"sortedPrefix"`
}

// SystemConfig global config
var SysConfig *Config

func init() {
	var err error
	SysConfig, err = Load("config.ini")
	if err != nil {
		log.Println(err)
	}
}

// Load get config from ini
func Load(configFileMame string) (*Config, error) {
	cfg, err := ini.Load(configFileMame)
	if err != nil {
		return nil, err
	}
	config := new(Config)
	err = cfg.MapTo(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
