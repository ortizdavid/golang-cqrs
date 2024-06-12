package main

import (
	"net/http"
	"github.com/ortizdavid/golang-cqrs/config"
	userController "github.com/ortizdavid/golang-cqrs/pkg/users/controllers"
)

func main() {
	mux := http.NewServeMux()
	userController.RegisterRoutes(mux)
	http.ListenAndServe(config.ListenAddr(), mux)
}