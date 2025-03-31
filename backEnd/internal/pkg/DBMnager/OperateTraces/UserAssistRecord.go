package OperateTraces

import (
	"gorm.io/gorm"
	"time"
)

// UserAssist活动记录表
type UserAssistRecord struct {
	LastExecuted   time.Time      `gorm:"type:timestamp;not null"`
	ExecutableName string         `gorm:"type:varchar(255);not null"`
	ExecutablePath string         `gorm:"type:text;not null"`
	RunCount       int            `gorm:"type:int;default:0;check:run_count >= 0"`
	FocusCount     int            `gorm:"type:int;default:0;check:focus_count >= 0"`
	CreateAt       time.Time      `gorm:"autoCreateTime"`
	DeleteAt       gorm.DeletedAt `gorm:"index"`
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	UUID           string         `gorm:"type:varchar(255);column:uuid"`
}

func (UserAssistRecord) TableName() string {
	return "user_assist_record"
}
