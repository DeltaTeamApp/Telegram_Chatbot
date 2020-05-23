package controller

import (
	"DeltaTeleBot/config"
	"DeltaTeleBot/handler/ggsheet"
	"fmt"
	"strconv"
	"time"
)

//GenSKU Gen SKU from given spreadsheet id
func GenSKU(arg string) {
	var msg string
	var markUpdate []string
	msgChan <- "GenSKU START"
	num, err := strconv.Atoi(arg)
	if err != nil {
		msgChan <- "Not a number"
		msgChan <- "exit"
		return
	}
	if num--; num < 0 {
		msgChan <- "Not a number"
		msgChan <- "exit"
		return
	}

	skuCfObj := config.GetSKUConfigObj()

	SKUResult := ggsheet.GetDataFromRage(skuCfObj.SheetID, skuCfObj.Table, skuCfObj.SKUCol, skuCfObj.MarkRow, skuCfObj.SKUCol, skuCfObj.MarkRow+num)
	timeData := time.Now()
	for i := 0; i < len(SKUResult); i++ {
		msg = msg + fmt.Sprintf("%+v\n", SKUResult[i])
		markUpdate = append(markUpdate, timeData.String())
	}
	msgChan <- msg

	err = ggsheet.UpdateDataInRange(skuCfObj.SheetID, skuCfObj.Table, skuCfObj.MarkCol, skuCfObj.MarkRow, skuCfObj.MarkCol, skuCfObj.MarkRow+num, markUpdate)
	if err != nil {
		msgChan <- fmt.Sprintf("Can not update SKU mark column : %+v:%+v, \nPlease do it manually : %+v", skuCfObj.MarkRow, skuCfObj.MarkRow+num, err)
	}

	err = skuCfObj.UpdateMarkRow(num)
	if err != nil {
		msgChan <- fmt.Sprintf("Can not update SKU mark row config : %+v:%+v, \nPlease do it manually : %+v", skuCfObj.MarkRow, skuCfObj.MarkRow+num, err)
	}
	msgChan <- "exit"
	return
}
