package DBMnager

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEntry     = gorm.ErrDuplicatedKey
	ErrUserNameOrPassword = errors.New("user name or password error")
)
