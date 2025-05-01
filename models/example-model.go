package models

type City struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Country    string  `json:"country"`
	Population int     `json:"population"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
