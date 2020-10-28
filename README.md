# google-my-business-go

**DO NOT USE THIS**
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
// Print all Locations under account.
func main() {
	var (
		clientID     = "<GOOGLEAPI_CLIENT_ID>"
		clientSecret = "<GOOGLEAPI_CLIENT_SECRET>"
		refreshToken = "<GOOGLEAPI_REFRESH_TOKEN>"
	)

	client, _ := New(clientID, clientSecret, refreshToken)
	accounts, _ := client.AccountAccess().List(url.Values{})

	var locations []*Location
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

