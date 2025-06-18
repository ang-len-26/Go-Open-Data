package models

type City struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Population int64   `json:"population"`
	CountryID  int     `json:"country_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
