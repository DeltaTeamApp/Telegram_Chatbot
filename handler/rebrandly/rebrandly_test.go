package rebrandly

import (
	"DeltaTeleBot/config"
	"testing"
)

func TestCreateShortLink(t *testing.T) {
	config.Init()
	var inputLink = []string{"http://deltastoreus.com/products/fm-fm0053v/", "http://deltastoreus.com/products/fm-fm0052w/", "http://deltastoreus.com/products/fm-fm0052o/", "http://deltastoreus.com/products/tumbler-t0054q1/"}
	var inputSlashtag = []string{"12052wf3", "12052of", "1205q1t", "1205q2t"}
	_, _, val := CreateShortLink(inputLink, inputSlashtag)
	if val != 4 {
		t.Errorf("Expect 4 - Receive : %+v", val)
	}
}
