package controller

import (
	"context"
	"time"

	"github.com/41x3n/TeleUtil/domain"
)

type HelpController struct {
	userRepository domain.UserRepository
}

func NewHelpController(ur domain.UserRepository) *HelpController {
	return &HelpController{
		userRepository: ur,
	}
}

func (sc *HelpController) Run(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := sc.userRepository.GetOrCreateByUserTelegramID(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
