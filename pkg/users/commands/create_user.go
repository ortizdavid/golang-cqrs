package commands

import (
	"time"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-cqrs/entities"
	"github.com/ortizdavid/golang-cqrs/repositories"
)

type CreateUserCommand struct {
	RoleId   	int
	UserName 	string
	Password 	string
}

type CreateUserCommandHandler struct {
	repository repositories.UserRepository
}

func (h CreateUserCommandHandler) Handle(cmd CreateUserCommand) (entities.User, error) {
	user := entities.User{
		UserId:    0,
		RoleId:    cmd.RoleId,
		UserName:  cmd.UserName,
		Password:  encryption.HashPassword(cmd.Password),
		IsActive:  true,
		Image:     "",
		UniqueId:  encryption.GenerateUUID(),
		Token:     encryption.GenerateRandomToken(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := h.repository.Create(user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}