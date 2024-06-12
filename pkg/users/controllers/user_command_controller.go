package controllers

import (
	"net/http"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/serialization"
	"github.com/ortizdavid/golang-cqrs/pkg/users/commands"
)

type UserCommandController struct {
}

func (us UserCommandController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var command commands.CreateUserCommand
	var handler commands.CreateUserCommandHandler

	err := serialization.DecodeJson(r.Body, &command)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
	}
	user, err := handler.Handle(command)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJsonSimple(w, http.StatusCreated, user)
}


func (us UserCommandController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var command commands.UpdateUserCommand
	var handler commands.UpdateUserCommandHandler

	id := r.FormValue("id")
	userId := conversion.StringToInt(id)

	err := serialization.DecodeJson(r.Body, &command)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
	}
	user, err := handler.Handle(command, userId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJsonSimple(w, http.StatusOK, user)
}


func (us UserCommandController) ActivateUser(w http.ResponseWriter, r *http.Request) {
	var command commands.ChangeUserStatusCommand
	var handler commands.ActivateUserCommandHandler

	err := serialization.DecodeJson(r.Body, &command)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
	}
	err = handler.Handle(command)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJsonSimple(w, http.StatusOK, nil)
}


func (us UserCommandController) DeactivateUser(w http.ResponseWriter, r *http.Request) {
	var command commands.ChangeUserStatusCommand
	var handler commands.DeactivateUserCommandHandler

	err := serialization.DecodeJson(r.Body, &command)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
	}
	err = handler.Handle(command)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJsonSimple(w, http.StatusOK, nil)
}