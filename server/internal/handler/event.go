package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wedding_api/internal/datastruct"
	"wedding_api/internal/service"

	"github.com/gorilla/mux"
)

func (h *Handler) EventHandler(w http.ResponseWriter, r *http.Request) {
	var status int = http.StatusOK
	var result string

	switch r.Method {
	case "GET":
		result, status = getEvent(h.eventService, r)
		break
	default:
		result = fmt.Sprintf("Bad method %s", r.Method)
		status = http.StatusBadGateway
		break
	}

	MakeHttpResponse(w, result, status)
}

func getEvent(eventService service.EventService, r *http.Request) (string, int) {
	params := mux.Vars(r)
	eventId, _ := strconv.ParseInt(params["eventId"], 10, 64)

	event, err := eventService.GetEventById(eventId)

	if err != nil {
		status := http.StatusNotFound
		message := fmt.Sprintf("Event with ID %v could not be found.", eventId)
		return MakeHttpError(status, message)
	}

	return eventToJson(event)
}

func eventToJson(event *datastruct.Event) (string, int) {
	data, err := json.Marshal(event)
	if (err != nil) {
		return MakeHttpError(500, "Unable to process event data.")
	}

	return string(data), http.StatusOK
}