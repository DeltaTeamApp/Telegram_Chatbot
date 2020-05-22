package ggsheet

import (
	"DeltaTeleBot/config"
	"testing"
)

func TestGetDataFromRage(t *testing.T) {
	config.Init()
	ServiceGGSheetInit()
	ggCfObj := config.GetGgsConfigObj()
	val := GetDataFromRage(ggCfObj.LinkSheetID, ggCfObj.LinkTableName, "A", 1, "A", 1)

	if val[0] != "a" {
		t.Errorf("Expect a - receive : %+v", val)
	}
}

func TestUpdateDataInRange(t *testing.T) {
	config.Init()
	ServiceGGSheetInit()
	ggCfObj := config.GetGgsConfigObj()
	val := []string{"Test", "Test 1"}
	err := UpdateDataInRange(ggCfObj.LinkSheetID, ggCfObj.LinkTableName, "V", 5147, "V", 5148, val)
	if err != nil {
		t.Errorf("Can not update data on google sheet %+v", err)
	}
}
