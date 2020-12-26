package models

type User struct {
	base     `gorm:"embedded"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Event    []Event `json:"event";gorm:"many2many"`
	Friend   []User  `json:"friend;gorm:"many2many"`
}
