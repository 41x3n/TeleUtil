package domain

import (
	"context"

	"gorm.io/gorm"
)

const (
	TableUser = "users"
)

type User struct {
	gorm.Model
	UserID       int64  `gorm:"not null"`
	IsBot        bool   `gorm:"not null"`
	FirstName    string `gorm:"not null"`
	LastName     *string
	UserName     *string
	LanguageCode *string
	IsActive     bool `gorm:"default:true"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByUserName(c context.Context, username string) (User, error)
	GetByID(c context.Context, userID int) (User, error)
	GetOrCreateByUserID(c context.Context, user *User) (User, error)
}
