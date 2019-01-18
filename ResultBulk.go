package postcodesio

type ResultsBulk struct {
	Status int `json:"status,omitempty"`
	Result []struct {
		Query  string   `json:"query, omitempty"`
		Result []Result `json:"result,omitempty"`
	} `json:"result, omitempty"`
}
