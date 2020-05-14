package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	ggsConfig    *viper.Viper
	ggsConfigObj *GgsConfigObj
)

//GgsConfigObj contains config for google sheet API
type GgsConfigObj struct {
	Filename      string
	Path          string
	LinkSheetID   string
	LinkTableName string
}

func ggsConfigInit() {
	ggsConfig = viper.New()

	ggsConfig = viper.New()
	ggsConfig.SetConfigName("googlesheet_secret")
	ggsConfig.SetConfigType("env")

	ggsConfig.AddConfigPath("./config_file/")
	ggsConfig.AddConfigPath("../../config_file/")
	ggsConfig.AddConfigPath("../config_file/")

	if err := ggsConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func findSecretFilePath(fileName string) (string, error) {
	var path string
	var err error
	base, err := filepath.Abs("../config_file/")
	path = filepath.Join(base, fileName)
	if fileExist(path) {
		return path, err
	}
	base, err = filepath.Abs("./config_file/")
	path = filepath.Join(base, fileName)
	if fileExist(path) {
		return path, err
	}
	base, err = filepath.Abs("../../config_file/")
	path = filepath.Join(base, fileName)
	if fileExist(path) {
		return path, err
	}
	return "", err
}

func ggsConfigObjInit() {
	var path string
	filename := ggsConfig.GetString("FILENAME")
	path, err := findSecretFilePath(filename)
	if err != nil {
		log.Fatalf("ggsConfigObjInit - find secret file : %+v", err)
	}

	ggsConfigObj = &GgsConfigObj{
		Filename:      filename,
		Path:          path,
		LinkSheetID:   ggsConfig.GetString("LINKSPREADSHEETID"),
		LinkTableName: ggsConfig.GetString("LINKTABLENAME"),
	}
}

//GetGgsConfigObj return config object for google sheeet
func GetGgsConfigObj() *GgsConfigObj {
	return ggsConfigObj
}
