package ggsheet

import (
	"DeltaTeleBot/config"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var (
	ggSheetService = new(sheets.Service)
)

//ServiceGGSheetInit create a service to google sheet
func ServiceGGSheetInit() {
	ggsConfig := config.GetGgsConfigObj()

	var err error
	data, err := ioutil.ReadFile(ggsConfig.Path)
	if err != nil {
		log.Fatalf("\nPkg: ggsheet - createClient - cannot read file : %+v\n", err)
	}

	ggConfig, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("\nPkg: ggsheet - createClient - cannot parse config : %+v\n", err)
	}

	client := ggConfig.Client(context.TODO())

	ggSheetService, err = sheets.New(client)
	if err != nil {
		log.Fatalf("\nPkg: ggsheet - createClient - cannot create service : %+v\n", err)
	}
}

//NewRange create range with input number and col
func parseRange(col1 string, firstNum int, col2 string, secondNum int) string {
	newStr := col1 + strconv.Itoa(firstNum) + ":" + col2 + strconv.Itoa(secondNum)
	return newStr
}

//GetDataFromRage get data from given range and sheetID
func GetDataFromRage(sheetID string, table string, col1 string, firstNum int, col2 string, secondNum int) (rows []string) {
	spreadsheetID := sheetID
	sheetRange := parseRange(col1, firstNum, col2, secondNum)
	readRange := table + "!" + sheetRange

	srv := ggSheetService
	//fetch data
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Printf("\nPkg: ggsheet - GetDataFromRage - Cannot get data: %v \n", err)
		rows = append(rows, "Pkg: ggsheet - GetDataFromRage - Cannot get data")
		return rows
	}

	if len(resp.Values) == 0 {
		log.Println("Pkg: ggsheet - GetDataFromRage - found nothing")
		rows = append(rows, "Pkg: ggsheet - GetDataFromRage - found nothing")
		return rows

	}

	for _, row := range resp.Values {
		rows = append(rows, row[0].(string))
		// log.Println(i, "  ", row[0])
	}
	return rows
}

//UpdateDataInRange Update data in given range with input value
func UpdateDataInRange(sheetID string, table string, col1 string, firstNum int, col2 string, secondNum int, val []string) error {
	sheetRange := parseRange(col1, firstNum, col2, secondNum)
	updateRange := table + "!" + sheetRange
	//create a service
	srv := ggSheetService
	//Create value to update
	var vr = new(sheets.ValueRange)
	for _, data := range val {
		vr.Values = append(vr.Values, []interface{}{data})
	}

	fmt.Println(vr)
	_, err := srv.Spreadsheets.Values.Update(sheetID, updateRange, vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Printf("\nPkg: ggsheet - UpdateDataInRange - Cannot update data: %v \n", err)
		return err
	}

	return nil
}
