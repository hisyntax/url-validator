# Domain
### You can use the tld package to get the list of all **TLDs** 

# Installation

1. You can use the below Go command to install the domain package
```sh
$ go get -u github.com/iqquee/tld
```
2. Import it in your code:
```sh
import "github.com/iqquee/tld"
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

	"github.com/iqquee/tld"
)

func main() {

	for _, v := range tld.Tlds {
		for _, tlds := range v.Tld {
			fmt.Printf("This is the tld response: %v\n", tlds.Tld)
		}
	}
}

```