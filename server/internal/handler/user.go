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

func (u *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	var status int = http.StatusOK
	var result string

	switch r.Method {
	case "GET":
		result, status = getUser(u.userService, r)
		break
	default:
		result = fmt.Sprintf("Bad method %s", r.Method)
		status = http.StatusBadGateway
		break
	}

	MakeHttpResponse(w, result, status)
}

func getUser(userService service.UserService, r *http.Request) (string, int) {
	params := mux.Vars(r)
	userId, _ := strconv.ParseInt(params["userId"], 10, 64)

	user, err := userService.GetUserById(userId)

	if err != nil {
		return MakeHttpError(
			404, 
			fmt.Sprintf("User wuth ID %v could not be found.", userId),
		), http.StatusNotFound
	}

	return userToJson(user), http.StatusOK
}

func userToJson(user *datastruct.Person) string {
	data, err := json.Marshal(user)
	if (err != nil) {
		return MakeHttpError(500, "Unable to process user data.")
	}

	return string(data)
}