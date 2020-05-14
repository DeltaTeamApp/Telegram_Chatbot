package name

import (
	"DeltaTeleBot/config"
	"testing"
)

func TestCreateFwdLink(t *testing.T) {
	config.Init()
	var inputLinks = []string{"testlink.com/products/hoodie-h0054x/",
		"testlink.com/products/hoodie-h0054y/",
		"testlink.com/products/hoodie-h00550/",
		"testlink.com/products/fm-fm00557/",
		"testlink.com/products/fm-fm00558/"}

	var tempLinks = []string{
		"1305h0054xdelta",
		"1305h00550delta",
		"1305h0054ydelta",
		"1305fm00557delta",
		"1305fm00558delta"}

	_, _, val := CreateFwdLink(inputLinks, tempLinks)
	if val != 5 {
		t.Errorf("Expect 5 - Receive : %+v", val)
	}
}
