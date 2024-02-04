package controller

import (
	"context"
	"log"
	"time"

	"github.com/41x3n/TeleUtil/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PhotoController struct {
	photoRepository domain.PhotoRepository
}

func NewPhotoController(pr domain.PhotoRepository) *PhotoController {
	return &PhotoController{
		photoRepository: pr,
	}
}

func (pc *PhotoController) Run(update *tgbotapi.Update, user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	photos := update.Message.Photo

	for _, p := range photos {
		photo := domain.Photo{
			UserTelegramID: user.TelegramID,
			FileID:         p.FileID,
			FileSize:       p.FileSize,
		}

		err := pc.photoRepository.Create(ctx, &photo)
		if err != nil {
			log.Println("Error saving photo", err)
		}
		log.Println("Photo saved - ", photo.FileID)
	}
	return nil
}
