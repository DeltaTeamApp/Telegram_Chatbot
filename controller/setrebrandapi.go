package controller

import (
	"DeltaTeleBot/handler/rebrandly"
	"fmt"
)

//SetNewRebrandAPIKey check valid of API key and set new
func SetNewRebrandAPIKey(arg string) string {
	err := rebrandly.ChangeAPIKey(arg)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Change API key successfully : %d links left", 500-rebrandly.CountLink())
}
