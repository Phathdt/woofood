package userhdl

import (
	"context"

	"user_service/module/user/usermodel"
)

type LoginUserRepo interface {
	FindUser(ctx context.Context, data map[string]interface{}) (*usermodel.User, error)
	CreateUserToken(ctx context.Context, user *usermodel.User) (*string, error)
}

type loginUserHdl struct {
	repo LoginUserRepo
}

func NewLoginUserHdl(repo LoginUserRepo) *loginUserHdl {
	return &loginUserHdl{repo: repo}
}

func (h *loginUserHdl) Response(ctx context.Context, data *usermodel.UserLogin) (*string, error) {
	user, err := h.repo.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(data.Password); err != nil {
		return nil, err
	}

	token, err := h.repo.CreateUserToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return token, nil
}
