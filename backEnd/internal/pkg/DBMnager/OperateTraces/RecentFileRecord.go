package OperateTraces

import (
	"gorm.io/gorm"
	"time"
)

type RecentFileRecord struct {
	FileName         string         `gorm:"type:varchar(255);not null"`
	FilePath         string         `gorm:"type:text;not null"`
	CreationTime     time.Time      `gorm:"type:timestamp;not null"`
	ModificationTime time.Time      `gorm:"type:timestamp;not null"`
	TargetPath       string         `gorm:"type:text;not null"`
	CreateAt         time.Time      `gorm:"autoCreateTime"`
	DeleteAt         gorm.DeletedAt `gorm:"index"`
	ID               uint           `gorm:"primaryKey;autoIncrement"`
	UUID             string         `gorm:"type:varchar(255);column:uuid"`
}

func (RecentFileRecord) TableName() string {
	return "recent_file_record"
}
