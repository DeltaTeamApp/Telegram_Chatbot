package rebrandly

import (
	"DeltaTeleBot/config"
	"testing"
)

var (
	ShortLink = shortLink
)

func TestCreateShortLink(t *testing.T) {
	config.Init()
	var inputLink = []string{"1205fm005uqdelta.testlink.com",
		"1205fm005urdelta.testlink.com",
		"1205fm005updelta.testlink.com",
		"1205fm005uldelta.testlink.com",
		"testlink.com/products/fm-fm0053v/"}

	var inputSlashtag = []string{"1205uqf",
		"1205urf",
		"1205upf",
		"1205ulf",
		"12052wf3"}
	_, _, val := CreateShortLink(inputLink, inputSlashtag)
	if val != 5 {
		t.Errorf("Expect 5 - Receive : %+v", val)
	}
	// for result := range results {
	// 	t.Log(result)
	// }
}

func TestShortLink(t *testing.T) {
	config.Init()
	loadConfig()
	result, val := shortLink("deltastoreus.com/products/fm-fm0053v/", "12052wf3")
	if val != nil {
		t.Errorf("Expect rebrand.ly/2020051401 - Receive : %+v", result)
	}
}

func TestCountLink(t *testing.T) {
	config.Init()
	loadConfig()
	val := CountLink()
	if val == -1 {
		t.Errorf("Expect unint - Receive %+v", val)
	}
}
