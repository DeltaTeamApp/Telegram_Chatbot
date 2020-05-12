package ggsheet

import (
	"DeltaTeleBot/config"
	"testing"
)

func TestGetDataFromRage(t *testing.T) {
	config.Init()
	ggCfObj := config.GetGgsConfigObj()

	val := GetDataFromRage(ggCfObj.LinkSheetID, ggCfObj.LinkTableName,
		"A", 1, "A", 1)

	if val[0] != "a" {
		t.Errorf("Expect a - receive : %+v", val)
	}
}
