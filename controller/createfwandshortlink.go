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
	firstNum, secondNum, err := splitArg(arg)
	if err != nil {
		msg = "Invalid format"
		msgChan <- msg
		return
	}
	ggsConfigObj := config.GetGgsConfigObj()
	nameConfigObj := config.GetNameConfigObj()
	rbConfigObj := config.GetRBConfigObj()

	storeLinks := ggsheet.GetDataFromRage(ggsConfigObj.LinkSheetID, ggsConfigObj.LinkTableName, nameConfigObj.StoreLinkColumn, firstNum, nameConfigObj.StoreLinkColumn, secondNum)
	tempLinks := ggsheet.GetDataFromRage(ggsConfigObj.LinkSheetID, ggsConfigObj.LinkTableName, nameConfigObj.TempLinkColumn, firstNum, nameConfigObj.TempLinkColumn, secondNum)

	fwdresults, fwdsucCount, fwderrCount := name.CreateFwdLink(storeLinks, tempLinks)

	for result := range fwdresults {
		msg = msg + fmt.Sprintf("%+v\n", result)
	}

	msg = msg + fmt.Sprintf("Success : %+v\nError : %+v\n", fwdsucCount, fwderrCount)
	msgChan <- msg

	slashTag := ggsheet.GetDataFromRage(ggsConfigObj.LinkSheetID, ggsConfigObj.LinkTableName, rbConfigObj.SlashTagCol, firstNum, rbConfigObj.SlashTagCol, secondNum)
	shortResults, shortSucCount, errSucCount := rebrandly.CreateShortLink(fwdresults, slashTag)
	msg = ""
	for result := range shortResults {
		msg = msg + fmt.Sprintf("%+v\n", result)
	}

	msg = msg + fmt.Sprintf("Success : %+v\nError : %+v\n", shortSucCount, errSucCount)
	linksCount := rebrandly.CountLink()
	msg = msg + fmt.Sprintf("Create : %+v\nLeft : %+v", linksCount, 500-linksCount)
	msgChan <- msg

	msg = "exit"
	msgChan <- msg

	return
}
