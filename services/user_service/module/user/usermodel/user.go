package usermodel

import (
	"github.com/phathdt/go-sdk/sdkcm"
	"golang.org/x/crypto/bcrypt"
	"user_service/common"
)

type User struct {
	sdkcm.SQLModel `json:",inline"`
	Email          string `json:"email"`
	Password       string `json:"-"`
}

func (user *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return sdkcm.ErrCustom(err, common.ErrPasswordNotMatched)
	}

	return nil
}

func (User) TableName() string {
	return "users"
}
