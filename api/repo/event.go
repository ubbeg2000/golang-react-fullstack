package repo

import (
	"golang-api/api/models"

	"gorm.io/gorm"
)

type EventRepo struct {
	conn *gorm.DB
}

func (r *EventRepo) FindAll(limit int, page int) []models.Event {
	var events []models.Event
	r.conn.Limit(limit).Offset(limit * (page - 1)).Find(&events)
	return events
}