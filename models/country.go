package models

type Country struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Capital     string  `json:"capital"`
	Population  int64   `json:"population"`
	Area        float64 `json:"area"`
	RegionID    int     `json:"region_id"`
	SubregionID int     `json:"subregion_id"`
}

type CountryDetail struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Capital    string     `json:"capital"`
	Population int64      `json:"population"`
	Area       float64    `json:"area"`
	Region     string     `json:"region"`
	Subregion  string     `json:"subregion"`
	Languages  []Language `json:"languages"`
	Currencies []Currency `json:"currencies"`
	Cities     []City     `json:"cities"`
}
