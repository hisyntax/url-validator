# Domain
### You can use the domain package to get the list of all **TLDs** and also to validate urls 

# Installation

1. You can use the below Go command to install the domain package
```sh
$ go get -u github.com/iqquee/domain
```
2. Import it in your code:
```sh
import "github.com/iqquee/domain"
```

# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```

# Tlds List
To get the list of all TLDs

```go
package main

import (
	"fmt"

	"github.com/iqquee/domain/tld"
)

func main() {

	t := tld.Tlds

	for _, v := range t {
		for _, tlds := range v.Tld {
			fmt.Printf("This is the tld response: %v\n", tlds.Tld)
		}
	}
}

```

# Validate URL

```go
package main

import (
	"fmt"

	urlReq "github.com/iqquee/domain/url"
)

func main() {
    
    //Example 1 - Valid URL
	url := "https://google.com"
	res, err := urlReq.Validate(url)
	if err != nil {
		fmt.Printf("This is the error recieved: %v\n", err.Error())
		return
	}

	fmt.Printf("This is the resspose: %v", res)
    //Response is url since its valid and its a url

    //Example 2 - text
	url := "hello"
	res, err := urlReq.Validate(url)
	if err != nil {
		fmt.Printf("This is the error recieved: %v\n", err.Error())
		return
	}

	fmt.Printf("This is the resspose: %v\n", res)
    //Response is text 

    //Example 3 - inValid URL
	url := "https://google .com"
	res, err := urlReq.Validate(url)
	if err != nil {
		fmt.Printf("This is the error recieved: %v", err.Error())
            //Error is inavlid url 
		return
	}

	fmt.Printf("This is the resspose: %v", res)


}
```