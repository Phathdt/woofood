package usermodel

import "github.com/phathdt/go-sdk/sdkcm"

type User struct {
	sdkcm.SQLModel
	Name string `json:"name"`
}

func (User) TableName() string {
	return "users"
}
