package model

import (
	"errors"
)

// 自定义一些错误
var (
	ERROR_USER_NOTEXIST = errors.New("error: user don't exist!")
	ERROR_USER_EXIST    = errors.New("error: user already exists!")
	ERROR_USER_PWD      = errors.New("error: user password incorrect!")
)
