package userstorage

import (
	"context"

	"github.com/phathdt/go-sdk/sdkcm"
	"user_service/module/user/usermodel"

	"gorm.io/gorm"
)

type sqlStorage struct {
	db *gorm.DB
}

func NewSqlStorage(db *gorm.DB) *sqlStorage {
	return &sqlStorage{db: db}
}

func (s *sqlStorage) CreateUser(ctx context.Context, data *usermodel.CreateUser) (uint32, error) {
	if err := s.db.Create(&data).Error; err != nil {
		return 0, sdkcm.ErrDB(err)
	}

	return data.ID, nil
}

func (s *sqlStorage) GetUserByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	var data usermodel.User

	db := s.db.Table(usermodel.User{}.TableName())

	if len(moreKeys) > 0 {
		for _, k := range moreKeys {
			db = db.Preload(k)
		}
	}

	result := db.Where(cond).Limit(1).Find(&data)
	if result.Error != nil {
		return nil, sdkcm.ErrDB(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, sdkcm.ErrDataNotFound
	}

	return &data, nil
}

func (s *sqlStorage) CreateUserToken(ctx context.Context, data *usermodel.UserToken) error {
	if err := s.db.Create(&data).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (s *sqlStorage) GetUserTokenByCondition(ctx context.Context, cond map[string]interface{}) (*usermodel.UserToken, error) {
	var data usermodel.UserToken

	db := s.db.Table(usermodel.UserToken{}.TableName())

	result := db.Where(cond).Limit(1).Find(&data)
	if result.Error != nil {
		return nil, sdkcm.ErrDB(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, sdkcm.ErrDataNotFound
	}

	return &data, nil
}
