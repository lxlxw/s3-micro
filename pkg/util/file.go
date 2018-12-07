package util

import "os"

// Files path exists
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

// Gets project path
func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

// Gets config path
func GetConfigPath(confName string) string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\" + "conf\\" + confName
	} else if ostype == "linux" {
		path = path + "/" + "conf/" + confName
	}
	return path
}

// Updates config file
func UpdateConfigFile(fileContent, fileName string) error {

	pathFile := GetConfigPath(fileName)

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
