package handler

import (
	"encoding/json"
	"fmt"
	"log"
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

// Making HTTP responses

func MakeHttpResponse(w http.ResponseWriter, result string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprint(w, result)
}

type httpError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"errorMessage"`
}

func MakeHttpError(statusCode int, message string) string {
	e := httpError{statusCode, message}
	data, err := json.Marshal(e)

	if (err != nil) {
		log.Fatalln(err)
	}

	return string(data)
}
