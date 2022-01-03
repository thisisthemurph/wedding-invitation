package repository

import (
	"errors"
	"wedding_api/internal/datastruct"
)

type EventQuery interface {
	CreateEvent(event datastruct.Event) (*int64, error)
	GetEventById(eventId int64) (*datastruct.Event, error)
}

type eventQuery struct{}

func (e *eventQuery) CreateEvent(event datastruct.Event) (*int64, error) {
	var createdId int64 = 1
	return &createdId, nil
}

func (e *eventQuery) GetEventById(eventId int64) (*datastruct.Event, error) {
	mocks := []datastruct.Event{
		{
			ID: 1, 
			Name: "Wedding", 
			Description: "The wedding of Max & Dani.",
		},
	}

	for _, event := range mocks {
		if event.ID == eventId {
			return &event, nil
		}
	}

	return nil, errors.New("Could not find event with given ID.")
}