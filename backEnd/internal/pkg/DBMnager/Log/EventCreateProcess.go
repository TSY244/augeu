package Log

import (
	"gorm.io/gorm"
	"time"
)

type EventCreateProcess struct {
	CreateAt          time.Time      `gorm:"autoCreateTime"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
	ID                uint           `gorm:"primaryKey;autoIncrement"`
	UUID              string         `gorm:"type:varchar(255);column:uuid"`
	EventID           string         `gorm:"column:event_id"`
	CreateUser        string         `gorm:"column:Create_User"`
	CreateUserDomain  string         `gorm:"column:Create_User_Domain"`
	NewProcessName    string         `gorm:"column:NewProcessName"`
	ParentProcessName string         `gorm:"column:ParentProcessName"`
	CommandLine       string         `gorm:"column:CommandLine;type:text"`
	Description       string         `gorm:"type:text"`
}

func (EventCreateProcess) TableName() string {
	return "event_createprocess"
}
