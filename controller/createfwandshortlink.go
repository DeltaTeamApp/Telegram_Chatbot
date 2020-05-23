package controller

import (
	"DeltaTeleBot/config"
	"DeltaTeleBot/handler/ggsheet"
	"DeltaTeleBot/handler/name"
	"DeltaTeleBot/handler/rebrandly"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	msgChan = make(chan string)
)

func splitArg(inputRange string) (firstNum int, secondNum int, err error) {
	inputRange = strings.TrimSpace(inputRange)

	separateIndex := strings.Index(inputRange, ":")
	if separateIndex == -1 {
		return -1, -1, errors.New("inputRange : Can not find : in inputRange")
	}
	firstNum, err1 := strconv.Atoi(inputRange[:separateIndex])
	secondNum, err2 := strconv.Atoi(inputRange[separateIndex+1:])
	if err1 != nil || err2 != nil {
		log.Printf("inputRange : \n\t %+v \n\t %+v", err1, err2)
		return -1, -1, errors.New("inputRange : Atoi")
	}

	if firstNum > secondNum {
		return -1, -1, errors.New("Invalid range input")
	}
	return firstNum, secondNum, nil
}

//GetUpdateChan return update channel for msg
func GetUpdateChan() chan string {
	return msgChan
}

//CreateFwdAndShortLinks input is range of google sheet and return msg for create forward links
func CreateFwdAndShortLinks(arg string) {
	var msg string
	CreateFwdAndShortLinksTimer := time.Now()

	msgChan <- "CreateFwdAndShortLinks START"
	firstNum, secondNum, err := splitArg(arg)
	if err != nil {
		msg = "Invalid format"
		msgChan <- msg

		msg = "exit"
		msgChan <- msg
		return
	}

	shortLinkConfigObj := config.GetShortLinkObj()

	storeLinks := ggsheet.GetDataFromRage(shortLinkConfigObj.SheetID, shortLinkConfigObj.Table, shortLinkConfigObj.StoreLinkCol, firstNum, shortLinkConfigObj.StoreLinkCol, secondNum)
	tempLinks := ggsheet.GetDataFromRage(shortLinkConfigObj.SheetID, shortLinkConfigObj.Table, shortLinkConfigObj.TempLinkCol, firstNum, shortLinkConfigObj.TempLinkCol, secondNum)

	fwdresults, fwdsucCount, fwderrCount := name.CreateFwdLink(storeLinks, tempLinks)

	err = ggsheet.UpdateDataInRange(shortLinkConfigObj.SheetID, shortLinkConfigObj.Table, shortLinkConfigObj.MiddleLinkCol, firstNum, shortLinkConfigObj.MiddleLinkCol, secondNum, fwdresults)
	if err != nil {
		msg = err.Error()
		msgChan <- msg

		msg = ""

		for i := 0; i < len(fwdresults); i++ {
			msg = msg + fmt.Sprintf("%+v\n", fwdresults[i])
		}
		msgChan <- msg
	}

	msg = fmt.Sprintf("NAME.COM\nSuccess : %+v\nError : %+v\n", fwdsucCount, fwderrCount)
	msgChan <- msg

	slashTag := ggsheet.GetDataFromRage(shortLinkConfigObj.SheetID, shortLinkConfigObj.Table, shortLinkConfigObj.SlashTagCol, firstNum, shortLinkConfigObj.SlashTagCol, secondNum)
	shortResults, shortSucCount, errSucCount := rebrandly.CreateShortLink(fwdresults, slashTag)
	msg = ""

	err = ggsheet.UpdateDataInRange(shortLinkConfigObj.SheetID, shortLinkConfigObj.Table, shortLinkConfigObj.ShortLinkCol, firstNum, shortLinkConfigObj.ShortLinkCol, secondNum, shortResults)
	if err != nil {
		msg = err.Error()
		msgChan <- msg

		msg = ""

		for i := 0; i < len(shortResults); i++ {
			msg = msg + fmt.Sprintf("%+v\n", shortResults[i])
		}
		msgChan <- msg
	}

	msg = fmt.Sprintf("REBRANDLY\nSuccess : %+v\nError : %+v\n", shortSucCount, errSucCount)
	msgChan <- msg

	linksCount := rebrandly.CountLink()
	msg = fmt.Sprintf("REBRANDLY COUNT\nCreated : %+v\nLeft : %+v", linksCount, 500-linksCount)
	msgChan <- msg

	msg = fmt.Sprintf("CreateFwdAndShortLinks : %+v", time.Since(CreateFwdAndShortLinksTimer))
	msgChan <- msg

	msg = "exit"
	msgChan <- msg

	return
}
