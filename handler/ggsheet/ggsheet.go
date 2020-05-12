package ggsheet

import (
	"DeltaTeleBot/config"
	"context"
	"io/ioutil"
	"log"
	"strconv"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

//GetDataFromRage get data from given range and sheetID
func GetDataFromRage(sheetID string, table string, col1 string, firstNum int, col2 string, secondNum int) (rows []string) {
	spreadsheetID := sheetID
	sheetRange := parseRange(col1, firstNum, col2, secondNum)
	readRange := table + "!" + sheetRange
	//create a service
	srv, err := createGGService()
	if err != nil {
		log.Printf("\nPkg: ggsheet - GetDataFromRage - create srv: %+v \n", err)
		rows = append(rows, "Pkg: ggsheet - GetDataFromRage - create srv")
		return rows
	}

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

func createGGService() (*sheets.Service, error) {
	ggsConfig := config.GetGgsConfigObj()

	data, err := ioutil.ReadFile(ggsConfig.Path)
	if err != nil {
		log.Printf("\nPkg: ggsheet - createClient - cannot read file : %+v\n", err)
		return nil, err
	}

	ggConfig, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	if err != nil {
		log.Printf("\nPkg: ggsheet - createClient - cannot parse config : %+v\n", err)
		return nil, err
	}

	client := ggConfig.Client(context.TODO())

	srv, err := sheets.New(client)
	if err != nil {
		log.Printf("\nPkg: ggsheet - createClient - cannot create service : %+v\n", err)
		return nil, err
	}

	return srv, nil
}

//NewRange create range with input number and col
func parseRange(col1 string, firstNum int, col2 string, secondNum int) string {
	newStr := col1 + strconv.Itoa(firstNum) + ":" + col2 + strconv.Itoa(secondNum)
	return newStr
}
