package postcodesio

// ResultsBulk is an organisation structure to return a bulk result
//
// Peter Holt <peter.holt@dochq.co.uk>
type ResultsBulk struct {
	Status int `json:"status,omitempty"`
	Result []struct {
		Query  string   `json:"query,omitempty"`
		Result []Result `json:"result,omitempty"`
	} `json:"result,omitempty"`
}
