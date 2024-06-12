package commands

import (
	"time"
	"github.com/ortizdavid/golang-cqrs/entities"
	"github.com/ortizdavid/golang-cqrs/repositories"
)

type UpdateUserCommand struct {
	NewRoleId   int
	NewUserName string
}

type UpdateUserCommandHandler struct {
	repository repositories.UserRepository
}

func (h UpdateUserCommandHandler) Handle(cmd UpdateUserCommand, userId int) (entities.User, error) {
	user, err := h.repository.GetById(userId)
	if err != nil {
		return entities.User{}, err
	}
	user.UserId = userId
	user.RoleId = cmd.NewRoleId
	user.UserName = cmd.NewUserName
	user.UpdatedAt = time.Now()
	
	err = h.repository.Update(user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}