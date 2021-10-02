# go-gmbapi

**!!DO NOT USE THIS!!**

This library was created to represent the idea of implementation. 
It is not yet feature ready and should not be used.

## Requirements
`go 1.14~`

## Installation
```shell
$ go get github.com/micheam/go-gmbapi
```

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	gmbapi "github.com/micheam/go-gmbapi"
)

// The following environment variables must be set
//
// - GMB_CLIENT_ID      string
// - GMB_CLIENT_SECRET  string
// - GMB_REFRESH_TOKEN  string

func main() {
	ctx := context.Background()
	client, _ := gmbapi.New()
	accounts, _ := client.AccountAccess().List(ctx, url.Values{})

	for acc := range accounts {
		acc := acc

		locs, _ := client.LocationAccess(acc).List(ctx, url.Values{})
		for loc := range locs {
			b, _ := json.Marshal(loc)
			fmt.Println(string(b))
		}
	}
}
```

The following directory contains other use examples.

- [list accounts](example/list-accounts)
- [list locations](example/list-all-locations)

## License
TBD

## Author
TBD
