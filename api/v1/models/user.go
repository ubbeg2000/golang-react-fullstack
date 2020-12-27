package models

type User struct {
	Base
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	RoleID   uint64  `json:"-"`
	Role     Role    `json:"role" gorm:"foreignKey:RoleID"`
	Event    []Event `json:"event" gorm:"many2many:event_user"`
	Friend   []User  `json:"friend" gorm:"many2many:user_friend"`
}
