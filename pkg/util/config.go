package util

import (
	"bytes"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Ks3 Ks3Config
		As3 As3Config
	}

	Ks3Config struct {
		AccessKey string `toml:"accesskey"`
		Secretkey string `toml:"secretkey"`
		Region    string `toml:"region"`
		Endpoint  string `toml:"endpoint"`
	}

	As3Config struct {
		AccessKey string `toml:"accesskey"`
		Secretkey string `toml:"secretkey"`
		Region    string `toml:"region"`
		Endpoint  string `toml:"endpoint"`
	}
)

var configInfo Config
var confName string
var ostype = runtime.GOOS

// 格式化toml文件
func EncodeConfig(conf *Config) (string, error) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(conf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// 解析配置文件
func ParseConfigFile(filePath string) {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if _, err := toml.Decode(string(fd), &configInfo); err != nil {
		panic(err)
	}
}

// 获取配置文件信息
func GetConfig() Config {
	return configInfo
}

// 生效配置文件
func SetConfig() {
	goEnv := os.Getenv("GOENV")
	switch goEnv {
	case "dev":
		confName = "app_dev.conf"
	case "test":
		confName = "app_test.conf"
	case "online":
		confName = "app_online.conf"
	default:
		confName = "app_dev.conf"
	}
	configFile := GetConfigPath(confName)
	ParseConfigFile(configFile)
}
