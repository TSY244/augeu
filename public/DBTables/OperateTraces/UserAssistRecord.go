package OperateTraces

import "time"

// UserAssist活动记录表
type UserAssistRecord struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;type:serial"`
	LastExecuted   time.Time `gorm:"type:timestamp;not null"`
	ExecutableName string    `gorm:"type:varchar(255);not null"`
	ExecutablePath string    `gorm:"type:text;not null"`
	RunCount       int       `gorm:"type:int;default:0;check:run_count >= 0"`
	FocusCount     int       `gorm:"type:int;default:0;check:focus_count >= 0"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

func (UserAssistRecord) TableName() string {
	return "user_assist_record"
}
