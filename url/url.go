package url

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iqquee/domain/request"
	"github.com/iqquee/domain/tld"
)

// The validateUrl() returns the request data-type (url or text) for a successful request and an error (if the url is invalid).
func Validate(url string) (string, error) {
	//loop through to get a match with the url provided
	for _, v := range tld.Tlds {
		for _, i := range v.Tld {

			tld1 := fmt.Sprintf(" .%s", i.Tld)
			tld2 := fmt.Sprintf(".%s", i.Tld)
			if strings.Contains(url, tld1) {
				return "url", errors.New("invalid url")
			} else if strings.Contains(url, tld2) {
				// send a request to the url provided to confirm if valid or not
				valid, _, err := request.NewRequest("GET", url)
				if err != nil {
					return "", errors.New("request url invalid")
				}
				fmt.Printf("This is the request url response: %v\n", string(valid))
				// return response
				return "url", nil
			}

		}
	}
	return "text", nil
}
