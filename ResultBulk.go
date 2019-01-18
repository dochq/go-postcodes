package postcodesio

type ResultsBulk struct {
	Status int      `json:"status,omitempty"`
	Result []Result `json:"result,omitempty"`
}
