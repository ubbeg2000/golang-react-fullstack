package models

type Tag struct {
	base `gorm:"embedded"`
	Name string `json:"name"`
}
