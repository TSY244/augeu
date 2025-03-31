package Log

import (
	"gorm.io/gorm"
	"time"
)

type SecurityEvent struct {
	CreateAt         time.Time      `gorm:"autoCreateTime"`
	DeleteAt         gorm.DeletedAt `gorm:"index"`
	ID               uint           `gorm:"primaryKey;autoIncrement"`
	UUID             string         `gorm:"type:varchar(255);column:uuid"`
	EventID          string         `gorm:"column:event_id"`
	LevelDisplayName string         `gorm:"column:LevelDisplayName"`
	Description      string         `gorm:"type:text"`
}

func (SecurityEvent) TableName() string {
	return "security_event"
}
