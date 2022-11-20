package url

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/hisyntax/domain/request"
)

func getTld() ([]string, error) {
	method := "GET"
	requestUrl := "https://tld1.p.rapidapi.com/GetTlds"
	token := "2bd0a67b11msh806b3bf797619bap1d4a9cjsn231f72cf61ea"
	tldArr, _, err := request.NewRequest(method, requestUrl, token)
	if err != nil {
		return nil, err
	}
	var arr []string
	//unmarshal into an array of string
	if err := json.Unmarshal(tldArr, &arr); err != nil {
		return nil, errors.New("error unmershalling request")
	}

	return arr, nil
}

func ValidateURL(url string) error {
	//send a get request to get the list of all tld
	arr, err := getTld()
	if err != nil {
		fmt.Printf("This is the error from the request: %v", err)
		return err
	}
	//loop through to get a match with the url provided
	for _, v := range arr {
		tld := "." + v
		if strings.Contains(url, tld) {
			for i, value := range url {
				//compare the value tld with the url tld if same
				if string(value) == "." {
					for index, vv := range url {
						prev := i - 1
						if index == prev {
							if string(vv) == "" {
								return errors.New("invalid url")
							}
						}
					}
				}
			}
		}
	}

	//send a request to the url provided to confirm if valid or not
	valid, _, err := request.NewRequest("GET", url, "")
	if err != nil {
		return errors.New("request url invalid")
	}
	fmt.Printf("This is the request url response: %v\n", string(valid))
	// return response

	return nil
}
