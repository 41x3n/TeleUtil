package repository

import (
	"context"

	"github.com/41x3n/TeleUtil/domain"
	"gorm.io/gorm"
)

type photoRepository struct {
	database *gorm.DB
	table    string
}

func NewPhotoRepository(db *gorm.DB, table string) domain.PhotoRepository {
	return &photoRepository{
		database: db,
		table:    table,
	}
}

func (pr *photoRepository) Create(c context.Context, photo *domain.Photo) error {
	return pr.database.WithContext(c).Table(pr.table).Create(photo).Error
}

func (pr *photoRepository) FetchByUser(c context.Context, user *domain.User) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := pr.database.WithContext(c).Table(pr.table).Where("telegram_id = ?", user.TelegramID).Find(&photos).Error
	return photos, err
}

func (pr *photoRepository) GetByFileID(c context.Context, fileID string) (domain.Photo, error) {
	var photo domain.Photo
	err := pr.database.WithContext(c).Table(pr.table).Where("file_id = ?", fileID).First(&photo).Error
	return photo, err
}
