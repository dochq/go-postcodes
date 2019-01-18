package postcodesio

type ResultSingle struct {
	Status int    `json:"status,omitempty"`
	Result Result `json:"result,omitempty"`
}
