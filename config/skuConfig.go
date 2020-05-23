package config

import (
	"log"

	"github.com/spf13/viper"
)

//SKUConfigObj ...
type SKUConfigObj struct {
	SheetID string
	Table   string
	SKUCol  string
	MarkCol string
	MarkRow int
}

var (
	skuConfig    *viper.Viper
	skuConfigObj *SKUConfigObj
)

func skuConfigInit() {
	skuConfig = viper.New()

	skuConfig.SetConfigName("sku_secret")
	skuConfig.SetConfigType("env")

	skuConfig.AddConfigPath("./config_file/")
	skuConfig.AddConfigPath("../../config_file/")
	skuConfig.AddConfigPath("../config_file/")

	if err := skuConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func skuConfigObjInit() {
	skuConfigObj = &SKUConfigObj{
		SheetID: skuConfig.GetString("SheetID"),
		Table:   skuConfig.GetString("Table"),
		SKUCol:  skuConfig.GetString("SKUCol"),
		MarkCol: skuConfig.GetString("MarkCol"),
		MarkRow: skuConfig.GetInt("MarkRow"),
	}
}

//GetSKUConfigObj ...
func GetSKUConfigObj() *SKUConfigObj {
	return skuConfigObj
}

//UpdateMarkRow ...
func (sCfObj *SKUConfigObj) UpdateMarkRow(num int) error {
	var err error
	sCfObj.MarkRow += num + 1
	skuConfig.Set("MarkRow", sCfObj.MarkRow)
	err = skuConfig.WriteConfig()
	if err != nil {
		sCfObj.MarkRow -= (num + 1)
		skuConfig.Set("MarkRow", sCfObj.MarkRow)
		return err
	}
	return nil
}
