package models

type Subregion struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RegionID int    `json:"region_id"`
}
