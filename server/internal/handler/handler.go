package handler

import (
	"net/http"
	"wedding_api/internal/service"
)

type IHandler interface {
	UserHandler(http.ResponseWriter, *http.Request)
	EventHandler(http.ResponseWriter, *http.Request)
}

type Handler struct {
	userService service.UserService
	eventService service.EventService
}

func MakeHandlers(
	userService *service.UserService, 
	eventService *service.EventService) IHandler {
	
		return &Handler{
		userService: *userService,
		eventService: *eventService,
	}
}