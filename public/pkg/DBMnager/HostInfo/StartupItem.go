package HostInfo

import (
	"gorm.io/gorm"
	"time"
)

// 启动项表
type StartupItem struct {
	ItemName       string         `gorm:"primaryKey;type:varchar(255)"`
	ExecutablePath string         `gorm:"type:text;not null"`
	SignatureInfo  string         `gorm:"type:text"`
	CreateAt       time.Time      `gorm:"autoCreateTime"`
	DeleteAt       gorm.DeletedAt `gorm:"index"`
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	UUID           string         `gorm:"type:varchar(255);column:uuid"`
}

func (StartupItem) TableName() string {
	return "startup_items"
}
