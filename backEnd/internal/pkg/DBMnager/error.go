package DBMnager

import (
	"gorm.io/gorm"
)

var (
	ErrDuplicateEntry = gorm.ErrDuplicatedKey
)
