package postcodesio

// ResultCodes is a small structure used in API response
// This is a seperate structure to kepe thing DRY
//
// Peter Holt <peter.holt@dochq.co.uk>
type ResultCodes struct {
	AdminCountry  string `json:"admin_country,omitempty"`
	AdminDistrict string `json:"admin_district,omitempty"`
	AdminWard     string `json:"admin_ward,omitempty"`
	Parish        string `json:"parish,omitempty"`
	CCG           string `json:"ccg,omitempty"`
	NUTS          string `json:"nuts,omitempty"`
}
