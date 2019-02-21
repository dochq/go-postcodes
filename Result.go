package postcodesio

// Result contains the data about the postcode
// This is shared across Lookup and LookupBulk
//
// Peter Holt <peter.holt@dochq.co.uk>
type Result struct {
	Postcode                 string      `json:"postcode,omitempty"`
	Quality                  int         `json:"quality,omitempty"`
	Eastings                 int         `json:"eastings,omitempty"`
	Northings                int         `json:"northings,omitempty"`
	Country                  string      `json:"country,omitempty"`
	NHSHA                    string      `json:"nhs_ha,omitempty"`
	AdminCountry             string      `json:"admin_country,omitempty"`
	AdminDistrict            string      `json:"admin_district,omitempty"`
	AdminWard                string      `json:"admin_ward,omitempty"`
	Longitude                float64     `json:"longitude,omitempty"`
	Latitude                 float64     `json:"latitude,omitempty"`
	ParlimentaryConstituency string      `json:"parlimentary_constituency,omitempty"`
	EuropeanElectoralRegion  string      `json:"european_electoral_region,omitempty"`
	PrimaryCareTrust         string      `json:"primary_care_trust,omitempty"`
	Region                   string      `json:"region,omitempty"`
	Parish                   string      `json:"parish,omitempty"`
	LSOA                     string      `json:"lsoa,omitempty"`
	MSOA                     string      `json:"msoa,omitempty"`
	CED                      string      `json:"ced,omitempty"`
	CCG                      string      `json:"ccg,omitempty"`
	NUTS                     string      `json:"nuts,omitempty"`
	Codes                    ResultCodes `json:"codes,omitempty"`
}
