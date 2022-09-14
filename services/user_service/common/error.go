package common

import "github.com/phathdt/go-sdk/sdkcm"

var (
	ErrPasswordNotMatched = sdkcm.CustomError("ErrPasswordNotMatched", "password not matched")
)
