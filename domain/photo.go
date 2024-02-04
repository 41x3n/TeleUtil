package domain

import (
	"context"

	"gorm.io/gorm"
)

const (
	TablePhoto = "photos"
)

type Photo struct {
	gorm.Model
	UserTelegramID int64  `gorm:"not null"`
	FileID         string `gorm:"not null"`
	FileSize       int    `gorm:"not null"`
}

type PhotoRepository interface {
	Create(c context.Context, photo *Photo) error
	FetchByUser(c context.Context, user *User) ([]Photo, error)
	GetByFileID(c context.Context, fileID string) (Photo, error)
}
