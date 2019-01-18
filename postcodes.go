package postcodesio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	ApiBaseUrl = "https://api.postcodes.io"
)

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

func Lookup(code string) (Result, error) {
	var res Result
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

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, err
	}

	return res, nil
}
