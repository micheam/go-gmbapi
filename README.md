# google-my-business-go

**!!DO NOT USE THIS!!**

This library was created to represent the idea of implementation. 
It is not yet feature ready and should not be used.

## Requirements
`go 1.14~`

## Installation
```shell
$ go get github.com/micheam/google-my-business-go
```

## Usage

```go
package main

import os
import github.com/micheam/google-my-business-go

// Print all Locations under account.
func init() {
	os.Setenv("GMB_CLIENT_ID", "THIS_IS_YOUR_CLIENT_ID")
	os.Setenv("GMB_CLIENT_SECRET", "THIS_IS_YOUR_CLIENT_SECRET")
	os.Setenv("GMB_REFRESH_TOKEN", "THIS_IS_YOUR_REFRESH_TOKEN")
}

func main() {
	client, _ := gmbapi.New()
	accounts, _ := client.AccountAccess().List(url.Values{})

	var locations []*gmbapi.Location
	for _, acc := range accounts.Accounts {
		acc := acc
		locs, _ := client.LocationAccess(acc).List(url.Values{})
		locations = append(locations, locs.Locations...)
	}
	b, _ := json.Marshal(locations)
	fmt.Println(string(b))
}
```

## License
TBD

## Author
TBD

