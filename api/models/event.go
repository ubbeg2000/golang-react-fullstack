package models

type Event struct {
	base        `gorm:"embedded"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tag         []Tag    `json:"tag";gorm:"many2many"`
	Location    Location `json:"location"`
	Author      User     `json:"author"`
	Attendee    []User   `json:"atendee"`
}
