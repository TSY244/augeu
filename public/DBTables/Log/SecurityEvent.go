package Log

import "time"

type SecurityEvent struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime       time.Time `gorm:"type:timestamp"`
	EventID          string    `gorm:"column:event_id"`
	LevelDisplayName string    `gorm:"column:LevelDisplayName"`
	Description      string    `gorm:"type:text"`
}

func (SecurityEvent) TableName() string {
	return "security_event"
}
