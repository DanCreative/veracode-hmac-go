# Veracode HMAC Authentication fo GO

> [!IMPORTANT]
> This package has moved into [DanCreative/veracode-go](https://github.com/DanCreative/veracode-go).


This project is a Go version of the Veracode API Signing library provided in the [Veracode Documentation](https://docs.veracode.com/r/c_hmac_signing_example_python). This is an updated version of someone else's project([antfie](https://github.com/antfie/veracode-go-hmac-authentication)), that caters for changes that Veracode have made more recently.

## Installation
```
go get -u github.com/DanCreative/veracode-hmac-go
```

## Example
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/DanCreative/veracode-hmac-go/vmac"
)

func main() {
	client := &http.Client{}
	method := "GET"
	apiUrl := "https://api.veracode.com/api/authn/v2/users/self"
	parsedUrl, _ := url.Parse(apiUrl)
	apiKey := "<VERACODE API KEY>"
	apiSecret := "<VERACODE API SECRET>"
	req, err := http.NewRequest(method, apiUrl, nil)

	if err != nil {
		panic(err)
	}
	authHeader, err := vmac.CalculateAuthorizationHeader(parsedUrl, method, apiKey, apiSecret)

	req.Header.Add("Authorization", authHeader)

	resp, _ := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Print(string(body[:]))
}
```
