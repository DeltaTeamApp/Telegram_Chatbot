package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	nameConfig    *viper.Viper
	nameConfigObj *NameConfigObject
)

//NameConfigObject contains config for name.com
type NameConfigObject struct {
	APIKey          string
	Domain          string
	Username        string
	StoreLinkColumn string
	TempLinkColumn  string
}

func nameConfigInit() {
	nameConfig = viper.New()
	nameConfig.SetConfigName("namedotcom_secret")
	nameConfig.SetConfigType("env")

	nameConfig.AddConfigPath(".")
	nameConfig.AddConfigPath("../../config_file/")
	nameConfig.AddConfigPath("../config_file/")

	if err := nameConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func nameConfigObjInit() {
	nameConfigObj = &NameConfigObject{
		APIKey:          nameConfig.GetString("NAMEAPIKEY"),
		Domain:          nameConfig.GetString("NAMEDOMAIN"),
		Username:        nameConfig.GetString("NAMEUSRNAME"),
		StoreLinkColumn: nameConfig.GetString("STORELINKCOLUMN"),
		TempLinkColumn:  nameConfig.GetString("TEMPFORWARDLINKCOLUMN"),
	}
}

//GetNameConfigObj return config object for name.com
func GetNameConfigObj() *NameConfigObject {
	return nameConfigObj
}
