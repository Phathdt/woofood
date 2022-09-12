package userhdl

import (
	"context"

	"user_service/module/user/usermodel"
)

type ListUserRepo interface {
	ListPost(ctx context.Context) ([]usermodel.User, error)
}

type listUserHdl struct {
	repo ListUserRepo
}

func NewListUserHdl(repo ListUserRepo) *listUserHdl {
	return &listUserHdl{repo: repo}
}

func (h *listUserHdl) Response(ctx context.Context) ([]usermodel.User, error) {
	return h.repo.ListPost(ctx)
}
