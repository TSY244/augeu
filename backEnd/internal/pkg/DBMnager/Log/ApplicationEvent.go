package Log

import (
	"gorm.io/gorm"
	"time"
)

type ApplicationEvent struct {
	CreateAt         time.Time      `gorm:"autoCreateTime"`
	DeleteAt         gorm.DeletedAt `gorm:"index"`
	ID               uint           `gorm:"primaryKey;autoIncrement"`
	UUID             string         `gorm:"type:varchar(255);column:uuid"`
	CreateTime       time.Time      `gorm:"type:timestamp"`
	EventID          string         `gorm:"column:event_id"`
	LevelDisplayName string         `gorm:"column:LevelDisplayName"`
	Description      string         `gorm:"type:text"`
}

func (ApplicationEvent) TableName() string {
	return "application_event"
}
