package usermodel

import (
	"github.com/phathdt/go-sdk/sdkcm"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	sdkcm.SQLModel `json:",inline"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
}

func (user *CreateUser) SetPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(hashedPassword)
}

func (CreateUser) TableName() string {
	return User{}.TableName()
}
