package rebrandly

import (
	"DeltaTeleBot/config"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
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

func countLink() int {
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
func CreateShortLink(inputLink []string, inputSlashTag []string) (result []string, successCount int, errorCount int) {
	if len(inputLink) != len(inputSlashTag) {
		log.Println("Pkg: rebrandly - CreateShortLink - Length not match")
		result = append(result, "Pkg: rebrandly - CreateShortLink - Length not match")
		return result, -1, -1
	}
	loadConfig()

	resultArr := make([]string, len(inputLink))
	steps := int(len(inputLink) / concurentNum)

	errChan := make(chan int8)
	defer close(errChan)

	var wg sync.WaitGroup
	wg.Add(concurentNum)
	for i := 0; i < concurentNum-1; i++ {
		go func(conI int) {
			for j := conI * steps; j < (conI+1)*steps; j++ {
				var err error
				resultArr[conI], err = shortLink(inputLink[conI], inputSlashTag[conI])
				if err != nil {
					errChan <- 1
				}
				errChan <- 0
			}
			wg.Done()
		}(i)
	}

	go func() {
		for i := concurentNum * steps; i < len(resultArr); i++ {
			var err error
			resultArr[i], err = shortLink(inputLink[i], inputSlashTag[i])
			if err != nil {
				errChan <- 1
			}
			errChan <- 0
			wg.Done()
		}
		errChan <- 2
	}()

	select {
	case sig := <-errChan:
		if sig == 0 {
			successCount++
		} else {
			if sig == 1 {
				errorCount++
			} else {
				break
			}
		}
	}

	wg.Wait()

	copy(result, resultArr)

	return result, successCount, errorCount
}
