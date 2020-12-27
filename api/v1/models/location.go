package models

type Location struct {
	Base
	Name       string `json:"name"`
	Country    string `json:"country"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Adress     string `json:"adress"`
	ZipCode    string `json:"zip_code"`
	LinkID     uint64 `json:"-"`
	Link       Link   `json:"link" gorm:"foreignKey:LinkID"`
	Image      string `json:"image"`
	Latitude   string `json:"latitude"`
	Longtitude string `json:"longitude"`
}
