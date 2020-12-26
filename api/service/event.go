package service

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/models"
	"github.com/ubbeg2000/golang-react-fullstack/api/repo"
)

type EventService struct {
	r repo.Repo
}

func (s *EventService) FindAll(limit int, page int) []models.Event {
	return s.r.Event.FindAll(limit, page)
}
