package name

import (
	"DeltaTeleBot/config"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	nameConfigObj *config.NameConfigObject
	concurentNum  = 4
)

func loadConfig() {
	nameConfigObj = new(config.NameConfigObject)
	nameConfigObj = config.GetNameConfigObj()
}

func forwardLink(link string, tempLink string) (string, error) {
	body := strings.NewReader(`{"host":"` + tempLink + `.` + nameConfigObj.Domain + `","forwardsTo":"` + link + `","type":"redirect"}`)

	req, err := http.NewRequest("POST", "https://api.name.com/v4/domains/"+nameConfigObj.Domain+"/url/forwarding", body)
	if err != nil {
		errLog := "Pkg: name - forwardLink - error : Create Request"
		log.Println(errLog)
		return errLog, err
	}
	req.SetBasicAuth(nameConfigObj.Username, nameConfigObj.APIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errLog := "Pkg: name - forwardLink - error : Make Request"
		log.Println(errLog)
		return errLog, err
	}
	resp.Body.Close()

	if resp.StatusCode == 200 {
		return (tempLink + nameConfigObj.Domain), nil
	}

	return ("error - " + strconv.Itoa(resp.StatusCode) + tempLink + nameConfigObj.Domain), errors.New("Pkg: name - forwardLink - sth went wrong")
}

//CreateFwdLink create forward links via name.com
func CreateFwdLink(inputLinks []string, tempLinks []string) (results []string, successCount int, errorCount int) {
	if len(inputLinks) != len(tempLinks) {
		log.Println("Pkg: name - CreateFwdLink - Length not match")
		results = append(results, "Pkg: name - CreateFwdLink - Length not match")
		return results, -1, -1
	}
	loadConfig()
	processLength := len(inputLinks)
	resultArr := make([]string, processLength)
	steps := int(processLength / concurentNum)
	errChan := make(chan int8)
	defer close(errChan)

	for i := 0; i < concurentNum; i++ {
		go func(id int) {
			for j := id * steps; j < (id+1)*steps; j++ {
				// log.Println("Process : ", j)
				var err error
				resultArr[j], err = forwardLink(inputLinks[j], tempLinks[j])
				if err == nil {
					errChan <- 0
				} else {
					errChan <- 1
				}
			}
		}(i)
	}

	go func() {
		for num := (concurentNum) * steps; num < processLength; num++ {
			// log.Println("Process : ", num)
			var err error
			resultArr[num], err = forwardLink(inputLinks[num], tempLinks[num])
			if err == nil {
				errChan <- 0
			} else {
				errChan <- 1
			}
		}
	}()

	for k := 0; k < processLength; k++ {
		// log.Println("chan : ", k)
		select {
		case sig := <-errChan:
			if sig == 0 {
				successCount++
			} else {
				errorCount++
			}
		}
	}

	copy(results, resultArr)
	return results, successCount, errorCount
}
