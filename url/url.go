package url

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/iqquee/domain/request"
)

var Token string

type Tld struct {
	Tlds []string
}

func getTld() (*Tld, error) {
	method := "GET"
	requestUrl := "https://tld1.p.rapidapi.com/GetTlds"
	token := Token
	tldArr, _, err := request.NewRequest(method, requestUrl, token)
	if err != nil {
		return nil, err
	}

	var res Tld
	//unmarshal into an array of string
	if err := json.Unmarshal(tldArr, &res.Tlds); err != nil {
		return nil, errors.New("error unmershalling request")
	}

	// fmt.Printf("This is the obj array byte: %v\n", string(tldArr))
	fmt.Printf("This is the obj array: %v\n", res.Tlds)

	return &res, nil
}

func ValidateURL(url string) (interface{}, string, error) {
	//send a get request to get the list of all tld
	arr, err := getTld()

	if err != nil {
		fmt.Printf("This is the error from the request: %v", err)
		return arr.Tlds, "", err
	}
	//loop through to get a match with the url provided
	for _, v := range arr.Tlds {
		tld1 := fmt.Sprintf(" .%s", v)
		tld2 := fmt.Sprintf(".%s", v)
		if strings.Contains(url, tld1) {
			return arr.Tlds, "url", errors.New("invalid url")
		} else if strings.Contains(url, tld2) {
			// send a request to the url provided to confirm if valid or not
			valid, _, err := request.NewRequest("GET", url, "")
			if err != nil {
				return arr.Tlds, "", errors.New("request url invalid")
			}
			fmt.Printf("This is the request url response: %v\n", string(valid))
			// return response
			return arr.Tlds, "url", nil
		}
	}
	return arr.Tlds, "text", nil
}
