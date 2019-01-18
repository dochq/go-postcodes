// go-postcodes is a golang package to ease the use, of the Postcodes.io API
// The name of the repo and the package differ since golang has isse with hyphens
// in package names
//
// Peter Holt <peter.holt@dochq.co.uk>
package postcodesio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Basic holder for API information
const (
	ApiBaseUrl = "https://api.postcodes.io"
)

var (
	// Package wide container for the HTTP client
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

// Function to perform a postcode lookup, A postcode is provided as a parameter
// the postcode is searched and the corrisponding retsult is sent back to the
// calling function
//
// Peter Holt <peter.holt@dochq.co.uk>
func Lookup(code string) (ResultSingle, error) {
	var res ResultSingle
	var err error

	if code == "" {
		err = fmt.Errorf("Postcode empty")
		return res, err
	}

	resp, err := httpClient.Get(ApiBaseUrl + "/postcodes/" + code)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return res, err
	}

	return res, nil
}
