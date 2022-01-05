package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"wedding_api/internal/datastruct"
	"wedding_api/internal/service"

	"github.com/gorilla/mux"
)

func (u *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	var status int = http.StatusOK
	var result string

	switch r.Method {
	case http.MethodGet:
		result, status = getUser(u.userService, r)
		break
	case http.MethodPost:
		result, status = createUser(u.userService, r)
		break
	default:
		result = fmt.Sprintf("Bad method %s", r.Method)
		status = http.StatusBadGateway
		break
	}

	MakeHttpResponse(w, result, status)
}

func createUser(userService service.UserService, r *http.Request) (string, int) {
	user, err := bodyToUser(r.Body)
	if err != nil {
		status := http.StatusNotFound
		return MakeHttpError(status, err.Error())
	}

	// Insert the user into the dtabase
	insertErr := userService.CreateUser(user)
	if insertErr != nil {
		return MakeHttpError(http.StatusBadRequest, err.Error())
	}

	return "", http.StatusCreated
}

func getUser(userService service.UserService, r *http.Request) (string, int) {
	params := mux.Vars(r)
	userId, _ := strconv.ParseInt(params["userId"], 10, 64)

	user, err := userService.GetUserById(userId)

	if err != nil {
		status := http.StatusNotFound
		message := fmt.Sprintf("User with ID %v could not be found.", userId)
		return MakeHttpError(status, message)
	}

	return personToJson(user)
}

// Converts a HTTP Body object into a Person object
func bodyToUser(body io.ReadCloser) (person datastruct.Person, err error) {
	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&person)
	
	return
}

// Converts a Person object into a JSON object
func personToJson(user *datastruct.Person) (string, int) {
	data, err := json.Marshal(user)
	if (err != nil) {
		return MakeHttpError(500, "Unable to process user data.")
	}

	return string(data), http.StatusOK
}