package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	orderservice "github.com/hermansetiawan77777/technical-warungpintar/internal/storage/order"
)

func Createorder(w http.ResponseWriter, r *http.Request) {
	type body struct {
		Username string `json:"username"`
		Id       int    `json:"id"`
		Name     string `json:"nameitem"`
		Price    int    `json:"price"`
	}
	var u *body
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
	if u.Username == "" {
		response["error"] = "Please fill username"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}

	payload := &orderservice.Bucket{
		Id:    u.Id,
		Name:  u.Name,
		Price: u.Price,
	}
	neworders := orderservice.Addorder(u.Username, payload)
	if neworders == nil {
		response["error"] = "Failed created new order, duplicated ID detected !"
		responseWithData(w, http.StatusInternalServerError, response)
		return
	}
	response["status"] = "success create new order !"
	response["result"] = neworders
	responseWithData(w, http.StatusOK, response)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	listorders := orderservice.Showorder(id)
	response := map[string]interface{}{}
	if listorders == nil {
		response["error"] = "Product Not Found"
		responseWithData(w, http.StatusNotFound, response)
		return
	}
	//var listresult []*orderservice.Order

	response["result"] = listorders

	responseWithData(w, http.StatusOK, response)
}

func responseWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
