package controllers

import "net/http"

func RegisterRoutes(router *http.ServeMux) {
	// commands
	router.HandleFunc("POST /api/users", UserCommandController{}.CreateUser)
	router.HandleFunc("PUT /api/users/{id}", UserCommandController{}.UpdateUser)
	router.HandleFunc("PUT /api/users/{id}/activate", UserCommandController{}.ActivateUser)
	router.HandleFunc("PUT /api/users/{id}/deactivate", UserCommandController{}.DeactivateUser)
	// queries
}