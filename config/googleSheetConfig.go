package config

import (
	"log"
	"runtime"

	"github.com/spf13/viper"
)

var (
	ggsConfig    *viper.Viper
	ggsConfigObj *GgsConfigObj
)

//GgsConfigObj contains config for google sheet API
type GgsConfigObj struct {
	Path          string
	LinkSheetID   string
	LinkTableName string
}

func ggsConfigInit() {
	ggsConfig = viper.New()

	ggsConfig = viper.New()
	ggsConfig.SetConfigName("googlesheet_secret")
	ggsConfig.SetConfigType("env")

	ggsConfig.AddConfigPath(".")
	ggsConfig.AddConfigPath("../../config_file/")
	ggsConfig.AddConfigPath("../config_file/")

	if err := ggsConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func ggsConfigObjInit() {
	var path string
	if runtime.GOOS == "windows" {
		path = ggsConfig.GetString("WINDOWSGGSSCRPATH")
	} else {
		path = ggsConfig.GetString("LINUXGGSSCRPATH")
	}

	ggsConfigObj = &GgsConfigObj{
		Path:          path,
		LinkSheetID:   ggsConfig.GetString("LINKSPREADSHEETID"),
		LinkTableName: ggsConfig.GetString("LINKTABLENAME"),
	}
}

//GetGgsConfigObj return config object for google sheeet
func GetGgsConfigObj() *GgsConfigObj {
	return ggsConfigObj
}
