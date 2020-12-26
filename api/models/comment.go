package models

type Comment struct {
	base   `gorm:"embedded"`
	Author User   `json:"author"`
	Body   string `json:"body"`
}
