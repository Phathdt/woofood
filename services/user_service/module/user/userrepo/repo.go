package userrepo

import (
	"context"

	"user_service/module/user/usermodel"
)

type UserStorage interface {
	ListUser(ctx context.Context) ([]usermodel.User, error)
}

type userRepo struct {
	s UserStorage
}

func NewUserRepo(s UserStorage) *userRepo {
	return &userRepo{s: s}
}

func (r *userRepo) ListPost(ctx context.Context) ([]usermodel.User, error) {
	return r.s.ListUser(ctx)
}
