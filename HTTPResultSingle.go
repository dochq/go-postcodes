package postcodesio

type HTTPResultSingle struct {
	Status int    `json:"status,omitempty"`
	Result Result `json:"result,omitempty"`
}
