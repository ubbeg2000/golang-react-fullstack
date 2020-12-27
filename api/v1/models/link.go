package models

type Link struct {
	Base
	ImageID uint64 `json:"-"`
	Image   Image  `json:"image" gorm:"foreignKey:ImageID"`
	Google  string `json:"google"`
}
