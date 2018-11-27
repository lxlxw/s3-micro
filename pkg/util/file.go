package util

import "os"

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取当前工作路径
func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

// 获取配置文件路径
func GetConfigPath(confName string) string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\" + "conf\\" + confName
	} else if ostype == "linux" {
		path = path + "/" + "conf/" + confName
	}
	return path
}

// 更新配置文件
func UpdateConfigFile(fileContent string) error {

	//TODO: update type conf
	pathFile := GetConfigPath("app_online.conf")

	exist, err := PathExists(pathFile)
	if err != nil {
		return err
	}
	if exist {
		err := os.Remove(pathFile)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(pathFile, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fileContent)
	return nil
}
