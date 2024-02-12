package controller

import (
	"context"
	"time"

	"github.com/41x3n/TeleUtil/domain"
)

type StartController struct {
	userRepository domain.UserRepository
}

func NewStartController(ur domain.UserRepository) *StartController {
	return &StartController{
		userRepository: ur,
	}
}

func (sc *StartController) Run(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := sc.userRepository.GetOrCreateByUserTelegramID(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
