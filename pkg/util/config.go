package util

import (
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

func GetConfig() Config {
	return configInfo
}

func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

func GetConfigPath(confName string) string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\" + "conf\\" + confName
	} else if ostype == "linux" {
		path = path + "/" + "conf/" + confName
	}
	return path
}

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
