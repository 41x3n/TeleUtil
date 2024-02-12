package repository

import (
	"context"

	"github.com/41x3n/TeleUtil/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
	table    string
}

func NewUserRepository(db *gorm.DB, table string) domain.UserRepository {
	return &userRepository{
		database: db,
		table:    table,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	return ur.database.WithContext(c).Table(ur.table).Create(user).Error
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	err := ur.database.WithContext(c).Table(ur.table).Find(&users).Error
	return users, err
}

func (ur *userRepository) GetByUserName(c context.Context, username string) (domain.User, error) {
	var user domain.User
	err := ur.database.WithContext(c).Table(ur.table).Where("user_name = ?", username).First(&user).Error
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, TelegramID int64) (domain.User, error) {
	var user domain.User
	err := ur.database.WithContext(c).Table(ur.table).Where("telegram_id = ?", TelegramID).First(&user).Error
	return user, err
}

func (ur *userRepository) GetOrCreateByUserTelegramID(c context.Context, user *domain.User) (domain.User, error) {
	err := ur.database.WithContext(c).Table(ur.table).Where("telegram_id = ?", user.TelegramID).FirstOrCreate(user).Error
	return *user, err
}
