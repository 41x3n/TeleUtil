package controller

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/domain"
	"github.com/41x3n/TeleUtil/rabbit"
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

func (pc *PhotoController) Run(update *tgbotapi.Update, user *domain.User, app *bootstrap.Application) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	photos := update.Message.Photo

	var pID *string

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
		pID = &photo.FileID
	}

	if pID == nil {
		return errors.New("no photo saved")
	}

	log.Println("Publishing message to RabbitMQ")

	err := rabbit.PublishMessage(app, *pID)
	if err != nil {
		log.Println("Error publishing message to RabbitMQ", err)
		return err
	}

	return nil
}
