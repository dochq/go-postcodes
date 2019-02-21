// Package postcodesio is a golang package to ease the use, of the Postcodes.io API
// The name of the repo and the package differ since golang has isse with hyphens
// in package names
//
// Peter Holt <peter.holt@dochq.co.uk>
package postcodesio

import (
	"crypto/tls"
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
	httpClient *http.Client
)

// init manages the disabling of TLS verification as this can cause issues on
// some K8 clusters, as the URL is hard coded, there is no issue with this
// since any other attempt at compromise would not map to struct
// I'd love to remove this, but all the time GKE instances miss loads of root CA's
// this has to be the solution for now.
//
// Peter Holt <peter.holt@dochq.co.uk>
func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient = &http.Client{Timeout: 10 * time.Second, Transport: tr}
}

// Lookup function to perform a postcode lookup, A postcode is provided as a parameter
// the postcode is searched and the corresponding retsult is sent back to the
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
