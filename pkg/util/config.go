package util

import (
	"bytes"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/lxlxw/toml"
)

type (
	Config struct {
		S3 S3Config
	}

	S3Config struct {
		AccessKey string `toml:"accesskey"`
		Secretkey string `toml:"secretkey"`
		Region    string `toml:"region"`
		Endpoint  string `toml:"endpoint"`
	}
)

var configInfo Config
var confName string
var ostype = runtime.GOOS

// Encode config
func EncodeConfig(conf *Config) (string, error) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(conf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Parse config file
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

// Gets config info
func GetConfig() Config {
	return configInfo
}

// Sets config info
func SetConfig(confName string) {
	configFile := GetConfigPath(confName)
	ParseConfigFile(configFile)
}
