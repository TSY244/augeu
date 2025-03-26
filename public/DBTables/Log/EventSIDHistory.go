package Log

import "time"

type EventLsassAccess struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime        time.Time `gorm:"type:timestamp"`
	EventID           string    `gorm:"column:event_id"`
	SourceName        string    `gorm:"column:Source_Name"`
	SourceDomain      string    `gorm:"column:Source_Domain"`
	SourceProcessName string    `gorm:"column:Source_Process_Name"`
	TargetProcessName string    `gorm:"column:Target_Process_Name"`
	Description       string    `gorm:"type:text"`
}

func (EventLsassAccess) TableName() string {
	return "event_lsassaccess"
}
