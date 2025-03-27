package Log

import (
	"gorm.io/gorm"
	"time"
)

type EventPowerShell struct {
	CreateAt    time.Time      `gorm:"autoCreateTime"`
	DeleteAt    gorm.DeletedAt `gorm:"index"`
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	UUID        string         `gorm:"type:varchar(255);column:uuid"`
	EventID     string         `gorm:"column:event_id"`
	Command     string         `gorm:"type:text"`
	Description string         `gorm:"type:text"`
}

func (EventPowerShell) TableName() string {
	return "event_powershell"
}
