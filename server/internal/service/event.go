package service

import (
	"wedding_api/internal/datastruct"
	"wedding_api/internal/repository"
)

type EventService interface {
	GetEventById(userId int64) (*datastruct.Event, error)
}

type eventService struct {
	dao repository.DAO
}

func NewEventService(dao repository.DAO) EventService {
	return &eventService{dao: dao}
}

func (e *eventService) GetEventById(eventId int64) (*datastruct.Event, error) {
	event, err := e.dao.NewEventQuery().GetEventById(eventId)
	// event, err := e.dao.NewEvetQuery().GetEventById(userId)

	if err != nil {
		return nil, err
	}

	ev := &datastruct.Event{
		ID: event.ID,
		Name: event.Name,
		Description: event.Description,
	}

	return ev, nil
}