package models

type Location struct {
	Name       string `json:"name"`
	Link       Link   `json:"link"`
	Image      string `json:"image"`
	Latitude   string `json:"latitude"`
	Longtitude string `json:"longitude"`
}
