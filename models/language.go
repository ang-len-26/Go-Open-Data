package models

type Language struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ISOCode    string `json:"iso_code"`
	NativeName string `json:"native_name"`
}
