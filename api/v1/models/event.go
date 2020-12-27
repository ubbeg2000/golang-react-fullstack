package models

type Event struct {
	Base
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tag         []Tag    `json:"tag" gorm:"many2many:event_tag"`
	LocationID  uint64   `json:"-"`
	Location    Location `json:"location" gorm:"foreignKey:LocationID"`
	AuthorID    uint64   `json:"-"`
	Author      User     `json:"author" gorm:"foreignKey:AuthorID"`
	Attendee    []User   `json:"atendee" gorm:"many2many:event_user"`
}
