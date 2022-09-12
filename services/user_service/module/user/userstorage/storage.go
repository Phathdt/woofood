package userstorage

import (
	"context"

	"user_service/module/user/usermodel"

	"github.com/phathdt/go-sdk/sdkcm"

	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *userStorage {
	return &userStorage{db: db}
}

func (s *userStorage) ListUser(ctx context.Context) ([]usermodel.User, error) {
	var data []usermodel.User

	db := s.db.Table(usermodel.User{}.TableName())

	if err := db.Find(&data).Error; err != nil {
		return nil, sdkcm.ErrDB(err)
	}

	return data, nil
}
