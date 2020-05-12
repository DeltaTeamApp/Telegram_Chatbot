package config

import (
	"testing"
)

var (
	FindSecretFilePath = findSecretFilePath
)

func TestGetTeleConfigObj(t *testing.T) {
	Init()
	obj := GetTeleConfigObj()
	val := obj.Debug
	if val != true {
		t.Errorf("Expect true - Receive %+v", val)
	}
}

func TestGetGgsConfigObj(t *testing.T) {
	Init()
	obj := GetGgsConfigObj()
	val := obj.LinkTableName
	if val != "Product" {
		t.Errorf("Expect Product - Receive %+v", val)
	}
}

func TestGetRBConfigObj(t *testing.T) {
	Init()
	obj := GetRBConfigObj()
	val := obj.SlashTagCol
	if val != "W" {
		t.Errorf("Expect W - Receive %+v", val)
	}
}

func TestGetNameConfigObj(t *testing.T) {
	Init()
	obj := GetNameConfigObj()
	val := obj.StoreLinkColumn
	if val != "T" {
		t.Errorf("Expect T - Receive %+v", val)
	}
}

func TestFindSecretFilePath(t *testing.T) {
	path, err := findSecretFilePath("googlesheet_template.env")
	if err != nil {
		t.Errorf("findSecretFilePath err : %+v - %+v", err, path)
	}
}
