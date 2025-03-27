package OperateTraces

import (
	"gorm.io/gorm"
	"time"
)

type PrefetchRecord struct {
	LaunchTime     time.Time      `gorm:"type:timestamp;not null"`
	ExecutableName string         `gorm:"type:varchar(255);not null"`
	ExecutableMD5  string         `gorm:"type:char(32);not null;column:executable_md5"`
	ExecutablePath string         `gorm:"type:text;not null"`
	CreateAt       time.Time      `gorm:"autoCreateTime"`
	DeleteAt       gorm.DeletedAt `gorm:"index"`
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	UUID           string         `gorm:"type:varchar(255);column:uuid"`
}

func (PrefetchRecord) TableName() string {
	return "prefetch_record"
}
