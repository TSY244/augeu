package Log

import (
	"gorm.io/gorm"
	"time"
)

type EventSIDHistory struct {
	CreateAt          time.Time      `gorm:"autoCreateTime"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
	ID                uint           `gorm:"primaryKey;autoIncrement"`
	UUID              string         `gorm:"type:varchar(255);column:uuid"`
	EventID           string         `gorm:"column:event_id"`
	SourceName        string         `gorm:"column:Source_Name"`
	SourceDomain      string         `gorm:"column:Source_Domain"`
	SourceProcessName string         `gorm:"column:Source_Process_Name"`
	TargetProcessName string         `gorm:"column:Target_Process_Name"`
	Description       string         `gorm:"type:text"`
}

func (EventSIDHistory) TableName() string {
	return "event_sid_history"
}
