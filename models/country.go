package models

type Country struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Capital    string  `json:"capital"`
	Region     string  `json:"region"`
	Subregion  string  `json:"subregion"`
	Population int     `json:"population"`
	Area       float64 `json:"area"`
}
