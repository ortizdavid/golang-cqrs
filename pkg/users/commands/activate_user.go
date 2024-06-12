package commands

import (
	"fmt"
	"time"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-cqrs/repositories"
)

type ChangeUserStatusCommand struct {
	UserId int
}

type ActivateUserCommandHandler struct {
	repository repositories.UserRepository
}

func (h ActivateUserCommandHandler) Handle(cmd ChangeUserStatusCommand) error {
	user, err := h.repository.GetById(cmd.UserId)
	if err != nil {
		return err
	}
	if user.IsActive {
		return fmt.Errorf("user already active")
	}
	user.IsActive = true
	user.Token = encryption.GenerateRandomToken()
	user.UpdatedAt = time.Now()
	return h.repository.Update(user)
}
