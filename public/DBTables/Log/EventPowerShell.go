package Log

import "time"

type EventPowerShell struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime  time.Time `gorm:"type:timestamp"`
	EventID     string    `gorm:"column:event_id"`
	Command     string    `gorm:"type:text"`
	Description string    `gorm:"type:text"`
}

func (EventPowerShell) TableName() string {
	return "event_powershell"
}
