package main

import (
	"fmt"
	"net/http"

	"github.com/hermansetiawan77777/technical-warungpintar/internal/httpserver"
)

func main() {
	appserver := &http.Server{
		//normally create Host in env so it will automatically update when the host is changed
		Addr:    "localhost:8080",
		Handler: httpserver.HandleRoutes(),
	}
	fmt.Println("Server start")
	appserver.ListenAndServe()

}
