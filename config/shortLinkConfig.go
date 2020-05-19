package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	shortLinkConfig *viper.Viper

	shortLinkConfigObj *ShortLinkConfigObj
)

//ShortLinkConfigObj ...
type ShortLinkConfigObj struct {
	SheetID      string
	Table        string
	StoreLinkCol string
	TempLinkCol  string
	SlashTagCol  string
}

func shortLinkConfigInit() {
	shortLinkConfig = viper.New()

	shortLinkConfig.SetConfigName("shortlink_secret")
	shortLinkConfig.SetConfigType("env")

	shortLinkConfig.AddConfigPath("./config_file/")
	shortLinkConfig.AddConfigPath("../../config_file/")
	shortLinkConfig.AddConfigPath("../config_file/")

	if err := shortLinkConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func shortLinkObjInit() {
	shortLinkConfigObj = &ShortLinkConfigObj{
		SheetID:      shortLinkConfig.GetString("SheetID"),
		Table:        shortLinkConfig.GetString("Table"),
		StoreLinkCol: shortLinkConfig.GetString("StoreLinkCol"),
		TempLinkCol:  shortLinkConfig.GetString("TempLinkCol"),
		SlashTagCol:  shortLinkConfig.GetString("SlashTag"),
	}
}

//GetShortLinkObj return short link config object
func GetShortLinkObj() *ShortLinkConfigObj {
	return shortLinkConfigObj
}
