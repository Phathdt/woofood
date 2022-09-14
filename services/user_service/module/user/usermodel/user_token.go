package usermodel

import "github.com/phathdt/go-sdk/sdkcm"

type UserToken struct {
	sdkcm.SQLModel `json:",inline"`
	UserID         uint32 `json:"user_id"`
	Token          string `json:"token"`
	Active         bool   `json:"active"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}
