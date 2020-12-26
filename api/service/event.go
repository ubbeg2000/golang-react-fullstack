package service

import (
	"golang-api/api/models"
	"golang-api/api/repo"
)

type EventService struct {
	r repo.Repo
}

func (s *EventService) FindAll(limit int, page int) []models.Event {
	return s.r.Event.FindAll(limit, page)
}
