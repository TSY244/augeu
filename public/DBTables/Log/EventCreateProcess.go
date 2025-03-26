package Log

import "time"

type EventCreateProcess struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime        time.Time `gorm:"type:timestamp"`
	EventID           string    `gorm:"column:event_id"`
	CreateUser        string    `gorm:"column:Create_User"`
	CreateUserDomain  string    `gorm:"column:Create_User_Domain"`
	NewProcessName    string    `gorm:"column:NewProcessName"`
	ParentProcessName string    `gorm:"column:ParentProcessName"`
	CommandLine       string    `gorm:"column:CommandLine;type:text"`
	Description       string    `gorm:"type:text"`
}

func (EventCreateProcess) TableName() string {
	return "event_createprocess"
}
