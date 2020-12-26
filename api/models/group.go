package models

type Group struct {
	base        `gorm:"embedded"`
	Name        string `json:"name"`
	Image       Image  `json:"image"`
	Description string `json:"description"`
	Member      []User `json:"member"`
}
