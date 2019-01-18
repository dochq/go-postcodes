package postcodesio

type Result struct {
	Postcode                 string
	Quality                  int
	Eastings                 int
	Northings                int
	Country                  string
	nhs_ha                   string
	AdminCountry             string
	AdminDistrict            string
	AdminWard                string
	Longitude                float64
	Latitude                 float64
	ParlimentaryConstituency string
	EuropeanElectoralRegion  string
	PrimaryCareTrust         string
	Region                   string
	Parish                   string
	LSOA                     string
	MSOA                     string
	CED                      string
	CCG                      string
	NUTS                     string
	Codes                    ResultCodes
}
