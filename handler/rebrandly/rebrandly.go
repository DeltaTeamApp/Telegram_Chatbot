package rebrandly

import (
	"DeltaTeleBot/config"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	rbConfig     *config.RBConfigObject
	concurentNum = 4
)

func loadConfig() {
	rbConfig = new(config.RBConfigObject)
	rbConfig = config.GetRBConfigObj()
}

func shortLink(forwardLink string, slashtag string) (result string, err error) {
	resp, err := http.Get("https://api.rebrandly.com/v1/links/new?apikey=" + rbConfig.APIKey + "&destination=http://" + forwardLink + "&slashtag=" + slashtag + "&domain[id]=" + rbConfig.DomainID)
	// fmt.Println("https://api.rebrandly.com/v1/links/new?apikey=" + apiKey + "&destination=http://" + forwardLinkSlice[i] + "&slashtag=" + slashTagSlice[i] + "&domain[id]=" + domainID)
	if err != nil {
		log.Println("Err shortLinkByRebrand")
		result = "Pkg: rebrandly - shortLink - create GET req error"
		return result, err
	}
	defer resp.Body.Close()
	//fmt.Println(forwardLinkSlice[i]+" => https://rebrand.ly/"+slashTagSlice[i], " : ", resp.StatusCode)
	if resp.StatusCode == 200 {
		result = "https://rebrand.ly/" + slashtag
		return result, nil
	}

	result = "error - " + strconv.Itoa(resp.StatusCode) + " : https://rebrand.ly/" + slashtag
	return result, errors.New("Pkg: rebrandly - shortLink - sth went wrong")
}

type linkCountType struct {
	Count int `json:"count"`
}

//CountLink show how many link has been create with env API key
func CountLink() int {
	//START : read number of shortlink created
	req, err := http.NewRequest("GET", "https://api.rebrandly.com/v1/links/count", nil)
	if err != nil {
		log.Println("Pkg: rebrandly - countLink - create GET req error")
		return -1
	}
	req.Header.Set("Apikey", rbConfig.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Pkg: rebrandly - countLink - call GET req error")
		return -1
	}
	defer resp.Body.Close()

	countByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Pkg: rebrandly - countLink - read resp error")
		return -1
	}

	var linkCounter linkCountType
	err = json.Unmarshal(countByte, &linkCounter)
	if err != nil {
		log.Println("Pkg: rebrandly - countLink - Unmarshal file error")
		return -1
	}
	//END : read number of shortlink created

	return linkCounter.Count
}

//CreateShortLink create short link with rebrandly
func CreateShortLink(inputLink []string, inputSlashTag []string) (results []string, successCount int, errorCount int) {
	if len(inputLink) != len(inputSlashTag) {
		log.Println("Pkg: rebrandly - CreateShortLink - Length not match")
		results = append(results, "Pkg: rebrandly - CreateShortLink - Length not match")
		return results, -1, -1
	}
	loadConfig()

	processLength := len(inputLink)

	resultArr := make([]string, processLength)
	// log.Printf("resultArr length : %+v \n", len(resultArr))
	steps := int(processLength / concurentNum)

	errChan := make(chan int8)
	defer close(errChan)

	for i := 0; i < concurentNum; i++ {
		go func(id int) {
			for j := id * steps; j < (id+1)*steps; j++ {
				// log.Println("Process : ", j)
				var err error
				resultArr[j], err = shortLink(inputLink[j], inputSlashTag[j])
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
			resultArr[num], err = shortLink(inputLink[num], inputSlashTag[num])
			if err == nil {
				errChan <- 0
			} else {
				errChan <- 1
			}
		}
	}()

	for k := 0; k < processLength; k++ {
		log.Println("chan : ", k)
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
