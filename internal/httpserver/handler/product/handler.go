package handler

import (
	"encoding/json"
	"net/http"

	productservice "github.com/hermansetiawan77777/technical-warungpintar/internal/storage/product"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	addProd := productservice.Addproduct()
	listproducts := productservice.Showproduct()
	response := map[string]interface{}{}
	if listproducts == nil {
		response["error"] = "Product Not Found"
		responseWithData(w, http.StatusNotFound, response)
		return
	}

	if !addProd {
		response["error"] = "Failed Generated Product"
		responseWithData(w, http.StatusNotFound, response)
		return
	}
	response["result"] = listproducts
	responseWithData(w, http.StatusOK, response)
}

func responseWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
