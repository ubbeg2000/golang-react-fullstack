package models

type Image struct {
	base   `gorm:"embedded"`
	URL    string `json:"url"`
	Name   string `json:"name"`
	Author User   `json:"author"`
}
