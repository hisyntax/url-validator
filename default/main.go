package main

import (
	"fmt"

	"github.com/iqquee/domain/tld"
	urlReq "github.com/iqquee/domain/url"
)

func main() {

	var counter int
	t := tld.Tlds
	for _, v := range t {
		counter += len(v.Tld)
	}

	fmt.Printf("This is the total tld number: %v\n", counter)
	url := "hello"
	res, err := urlReq.ValidateURL(url)
	if err != nil {
		fmt.Printf("This is the error received: %v\n", err.Error())
		return
	}

	fmt.Printf("This is the response: %v\n", res)

}
