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
	var res HTTPResultSingle
	var err error

	if code == "" {
		err = fmt.Errorf("Postcode empty")
		return res.Result, err
	}

	resp, err := httpClient.Get(ApiBaseUrl + "/postcodes/" + code)

	if err != nil {
		return res.Result, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return res.Result, err
	}

	return res.Result, nil
}
