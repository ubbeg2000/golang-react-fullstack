package models

type Group struct {
	Base
	Name        string `json:"name"`
	ImageID     uint64 `json:"-"`
	Image       Image  `json:"image" gorm:"foreignKey:ImageID"`
	Description string `json:"description"`
	Member      []User `json:"member" gorm:"many2many:group_user"`
}
