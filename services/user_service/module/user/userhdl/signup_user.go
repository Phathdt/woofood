package userhdl

import (
	"context"

	"user_service/module/user/usermodel"
)

type CreateUserRepo interface {
	CreateUser(ctx context.Context, data *usermodel.CreateUser) (uint32, error)
}

type createUserHdl struct {
	repo CreateUserRepo
}

func NewCreateUserHdl(repo CreateUserRepo) *createUserHdl {
	return &createUserHdl{repo: repo}
}

func (h *createUserHdl) Response(ctx context.Context, data *usermodel.CreateUser) error {
	_, err := h.repo.CreateUser(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
