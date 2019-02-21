package postcodesio

// Organisation structure to return a single result
//
// Peter Holt <peter.holt@dochq.co.uk>
type ResultSingle struct {
	Status int    `json:"status,omitempty"`
	Result Result `json:"result,omitempty"`
}
