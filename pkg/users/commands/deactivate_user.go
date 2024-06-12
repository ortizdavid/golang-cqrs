package commands

import (
	"fmt"
	"time"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-cqrs/repositories"
)

type DeactivateUserCommandHandler struct {
	repository repositories.UserRepository
}

func (h DeactivateUserCommandHandler) Handle(cmd ChangeUserStatusCommand) error {
	user, err := h.repository.GetById(cmd.UserId)
	if err != nil {
		return err
	}
	if !user.IsActive {
		return fmt.Errorf("user already inactive")
	}
	user.IsActive = false
	user.Token = encryption.GenerateRandomToken()
	user.UpdatedAt = time.Now()
	return h.repository.Update(user)
}