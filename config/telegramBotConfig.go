package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	// viper
	teleConfig *viper.Viper

	// object
	teleConfigObj *TeleConfigObject
)

//TeleConfigObject telegram bot config
type TeleConfigObject struct {
	APIKey       string
	UpdateOffset int
	TimeOut      int
	Debug        bool
}

func teleConfigInit() {
	teleConfig = viper.New()

	teleConfig.SetConfigName("telebot_secret")
	teleConfig.SetConfigType("env")

	teleConfig.AddConfigPath("./config_file/")
	teleConfig.AddConfigPath("../../config_file/")
	teleConfig.AddConfigPath("../config_file/")

	if err := teleConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func teleConfigObjInit() {
	var debug bool

	if teleConfig.GetInt("DEBUG") == 1 {
		debug = true
	} else {
		debug = false
	}

	teleConfigObj = &TeleConfigObject{
		APIKey:       teleConfig.GetString("TELEGRAMBOTAPIKEY"),
		TimeOut:      teleConfig.GetInt("TELETIMEOUT"),
		UpdateOffset: teleConfig.GetInt("TELEOFFSET"),
		Debug:        debug,
	}

}

//GetTeleConfigObj return object contains config for telegrambot
func GetTeleConfigObj() *TeleConfigObject {
	return teleConfigObj
}
