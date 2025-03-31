package HostInfo

import (
	"gorm.io/gorm"
	"time"
)

// 镜像劫持表
type ImageHijack struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	UUID         string         `gorm:"type:varchar(255);column:uuid"`
	CreateAt     time.Time      `gorm:"autoCreateTime"`
	DeleteAt     gorm.DeletedAt `gorm:"index"`
	HijackName   string         `gorm:"type:varchar(255)"`
	HijackPath   string         `gorm:"type:text;not null"`
	HijackStatus bool           `gorm:"not null;default:false"`
}

func (ImageHijack) TableName() string {
	return "image_hijacks"
}
