package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	rbConfig *viper.Viper

	rbConfigObj *RBConfigObject
)

//RBConfigObject config for rebrandly
type RBConfigObject struct {
	APIKey      string
	DomainID    string
	SlashTagCol string
}

//SetAPIKey set new API key for rebrandly and write back to config file
func (r *RBConfigObject) SetAPIKey(newAPIKey string) error {
	var err error
	bkAPIKey := r.APIKey

	r.APIKey = newAPIKey
	rbConfig.Set("REBRANDLYAPIKEY", newAPIKey)
	err = rbConfig.WriteConfig()
	if err != nil {
		r.APIKey = bkAPIKey
		rbConfig.Set("REBRANDLYAPIKEY", bkAPIKey)
	}

	return err
}

func rbConfigInit() {
	rbConfig = viper.New()
	rbConfig.SetConfigName("rb_secret")
	rbConfig.SetConfigType("env")

	rbConfig.AddConfigPath(".")
	rbConfig.AddConfigPath("../../config_file/")
	rbConfig.AddConfigPath("../config_file/")

	if err := rbConfig.ReadInConfig(); err != nil {
		log.Fatalf("Can not read the config file : %+v", err)
	}
}

func rbConfigObjInit() {
	rbConfigObj = &RBConfigObject{
		APIKey:      rbConfig.GetString("REBRANDLYAPIKEY"),
		DomainID:    rbConfig.GetString("REBRANDLYDOMAINID"),
		SlashTagCol: rbConfig.GetString("SLASHTAGCOLUMN"),
	}
}

//GetRBConfigObj return config object for Rebrandly
func GetRBConfigObj() *RBConfigObject {
	return rbConfigObj
}
