package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Service struct {
		AppMode    string `yaml:"app_mode"`
		ServerPort int    `yaml:"server_port"`
	}
	MySQL struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		DBName string `yaml:"dbname"`
	}
}

var AppConfig Config

func Load() {
	file, err := os.Open("/config/workspace/sources/golang/Pharmacy-POS/config.yaml")
	if err != nil {
		fmt.Println("打开配置文件失败")
		return
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		fmt.Println("读取配置文件内容失败")
		return
	}

	file.Close()
}
