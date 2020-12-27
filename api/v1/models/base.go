package models

type Base struct {
	ID        uint64 `json:"id" gorm:"foreignKey;autoIncrement"`
	CreatedAt uint   `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt uint   `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
