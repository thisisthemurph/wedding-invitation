package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"wedding_api/internal/datastruct"
	"wedding_api/internal/httputils"
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
	
	httputils.MakeHttpResponse(w, result, status)
}

func createUser(userService service.UserService, r *http.Request) (string, int) {
	user, err := bodyToUser(r.Body)
	if err != nil {
		status := http.StatusNotFound
		return httputils.MakeHttpError(status, err.Error())
	}

	// Insert the user into the dtabase
	person, insertErr := userService.CreateUser(user)
	if insertErr != nil {
		return httputils.MakeHttpError(http.StatusBadRequest, insertErr.Error())
	}

	json, _ := personToJson(person)
	return json, http.StatusCreated
}

func getUser(userService service.UserService, r *http.Request) (string, int) {
	params := mux.Vars(r)
	userId, _ := strconv.ParseInt(params["userId"], 10, 64)

	user, err := userService.GetUserById(userId)

	if err != nil {
		status := http.StatusNotFound
		message := fmt.Sprintf("User with ID %v could not be found.", userId)
		return httputils.MakeHttpError(status, message)
	}

	json, _ := personToJson(user)
	return json, http.StatusOK
}

// Converts a HTTP Body object into a Person object
func bodyToUser(body io.ReadCloser) (person datastruct.Person, err error) {
	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&person)
	
	return
}

// Converts a Person object into a JSON object
func personToJson(user *datastruct.Person) (string, error) {
	data, err := json.Marshal(user)
	if (err != nil) {
		return "", err
	}

	return string(data), nil
}