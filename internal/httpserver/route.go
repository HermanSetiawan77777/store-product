package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	handlerorder "github.com/hermansetiawan77777/technical-warungpintar/internal/httpserver/handler/order"
	handlerprod "github.com/hermansetiawan77777/technical-warungpintar/internal/httpserver/handler/product"
	handleruser "github.com/hermansetiawan77777/technical-warungpintar/internal/httpserver/handler/user"
)

func HandleRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/newuser", handleruser.CreateUser).Methods(http.MethodPost)

	//please run add function "NewUser" first before "LoginUser"
	r.HandleFunc("/login", handleruser.LoginUser).Methods(http.MethodPost)
	r.HandleFunc("/product", handlerprod.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/order", handlerorder.Createorder).Methods(http.MethodPost)
	r.HandleFunc("/order/{id}", handlerorder.GetOrder).Methods(http.MethodGet)
	return r
}
