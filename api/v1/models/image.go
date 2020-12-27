package models

type Image struct {
	Base
	URL      string `json:"url"`
	Name     string `json:"name"`
	AuthorID uint64 `json:"-"`
	Author   User   `json:"author" gorm:"foreignKey:AuthorID"`
}
