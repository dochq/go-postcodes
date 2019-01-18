package postcodesio

type HTTPResultsBulk struct {
	Status int      `json:"status,omitempty"`
	Result []Result `json:"result,omitempty"`
}
