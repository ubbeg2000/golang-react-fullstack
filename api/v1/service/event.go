package service

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/models"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/repo"
)

type EventService struct{ repo repo.Repo }

func (s *EventService) FindAll(limit int, page int) []models.Event {
	return s.repo.Event.FindAll(limit, page)
}
