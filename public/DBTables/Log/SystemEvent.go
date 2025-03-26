package Log

import "time"

type SystemEvent struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime       time.Time `gorm:"type:timestamp"`
	EventID          string    `gorm:"column:event_id"`
	LevelDisplayName string    `gorm:"column:LevelDisplayName"`
	Description      string    `gorm:"type:text"`
}

func (SystemEvent) TableName() string {
	return "system_event"
}
