package handler

import (
	"encoding/json"
	"io"
	"net/http"

	userservice "github.com/hermansetiawan77777/technical-warungpintar/internal/storage/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	type newUser struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		Password       string `json:"password"`
		RetypePassword string `json:"retypepassword"`
	}
	var u *newUser
	err := json.NewDecoder(r.Body).Decode(&u)
	response := map[string]interface{}{}
	if err != nil {
		if err == io.EOF {
			response["error"] = "Please fill payload"
			responseWithData(w, http.StatusBadRequest, response)
			return
		}
		response["error"] = "Error parsing payload"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}
	if u.Id == "" {
		response["error"] = "Please fill id"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}
	if u.Name == "" {
		response["error"] = "Please fill name"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}
	if u.Password == "" {
		response["error"] = "Please fill password"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}
	if u.RetypePassword == "" {
		response["error"] = "Please fill retype-password"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}
	if u.RetypePassword != u.Password {
		response["error"] = "password and retype-password not match !"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}

	newUsers := userservice.AddNewUser(u.Id, u.Name, u.Password)
	if newUsers == nil {
		response["error"] = "Failed created new user, duplicated ID detected !"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}

	response["status"] = "success create new user !"
	response["result"] = newUsers

	responseWithData(w, http.StatusOK, response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	type login struct {
		Id       string `json:"id"`
		Password string `json:"password"`
	}
	var u *login
	err := json.NewDecoder(r.Body).Decode(&u)
	response := map[string]interface{}{}
	if err == io.EOF {
		response["error"] = "Id Not Found !"
		responseWithData(w, http.StatusNotFound, response)
		return
	}
	ExistingUser := userservice.LoginUser(u.Id, u.Password)

	if ExistingUser == nil {
		response["error"] = "Wrong ID or Password !"
		responseWithData(w, http.StatusNotFound, response)
		return
	}
	response["id"] = ExistingUser.Id
	response["name"] = ExistingUser.Name
	response["password"] = ExistingUser.Password
	response["Token"] = ExistingUser.LoginToken
	responseWithData(w, http.StatusOK, response)
}

func responseWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
