package userrepo

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/phathdt/go-sdk/sdkcm"
	"user_service/module/user/usermodel"
)

type UserStorage interface {
	GetUserByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.CreateUser) (uint32, error)
	CreateUserToken(ctx context.Context, data *usermodel.UserToken) error
	GetUserTokenByCondition(ctx context.Context, cond map[string]interface{}) (*usermodel.UserToken, error)
}

type repo struct {
	store UserStorage
}

func NewUserRepo(s UserStorage) *repo {
	return &repo{store: s}
}

func (r *repo) CreateUser(ctx context.Context, data *usermodel.CreateUser) (uint32, error) {
	user := usermodel.CreateUser{
		SQLModel: *sdkcm.NewSQLModel(),
		Email:    data.Email,
		Password: data.SetPassword(data.Password),
	}

	return r.store.CreateUser(ctx, &user)
}

func (r *repo) FindUser(ctx context.Context, data map[string]interface{}) (*usermodel.User, error) {
	return r.store.GetUserByCondition(ctx, data)
}

func (r *repo) FindUserToken(ctx context.Context, data map[string]interface{}) (*usermodel.UserToken, error) {
	return r.store.GetUserTokenByCondition(ctx, data)
}

func (r *repo) CreateUserToken(ctx context.Context, user *usermodel.User) (*string, error) {
	key := sdkcm.RandStringBytes(24)
	claims := jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"key":   key,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	userToken := usermodel.UserToken{
		SQLModel: *sdkcm.NewSQLModel(),
		UserID:   user.ID,
		Token:    key,
		Active:   true,
	}

	if err = r.store.CreateUserToken(ctx, &userToken); err != nil {
		return nil, err
	}

	return &t, nil
}
