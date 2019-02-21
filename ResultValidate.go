package postcodesio

// ResultValidate placed for future work to be done
//
// Peter Holt <peter.holt@dochq.co.uk>
type ResultValidate struct {
	Status int  `json:"status,omitempty"`
	Result bool `json:"result,omitempty"`
}
