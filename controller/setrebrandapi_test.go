package controller

import (
	"DeltaTeleBot/config"
	"strings"
	"testing"
)

func TestSetNewRebrandAPIKey(t *testing.T) {
	config.Init()
	val := SetNewRebrandAPIKey("4bae7d3da30b47789ffcb69f4ded9e1e")
	if strings.Contains(val, "successfully") {
		t.Error("Can not change the API key")
	}
}
