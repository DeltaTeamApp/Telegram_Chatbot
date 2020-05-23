package config

import (
	"os"
)

//Init load data to object
func Init() {
	teleConfigInit()
	teleConfigObjInit()

	rbConfigInit()
	rbConfigObjInit()

	nameConfigInit()
	nameConfigObjInit()

	ggsConfigInit()
	ggsConfigObjInit()

	shortLinkConfigInit()
	shortLinkObjInit()

	skuConfigInit()
	skuConfigObjInit()
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
