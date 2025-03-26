package Log

import "time"

type Event4624 struct {
	ID                        uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime                time.Time `gorm:"type:timestamp"`
	EventID                   string    `gorm:"column:event_id"`
	SourceIP                  string    `gorm:"column:source_ip"`
	SourceName                string    `gorm:"column:source_name"`
	TargetName                string    `gorm:"column:target_name"`
	LogonType                 string    `gorm:"column:logon_type"`
	LogonProc                 string    `gorm:"column:logon_proc"`
	LogonProcessName          string    `gorm:"column:LogonProcessName"`
	AuthenticationPackageName string    `gorm:"column:AuthenticationPackageName"`
	Description               string    `gorm:"type:text"`
}

func (Event4624) TableName() string {
	return "event_4624"
}
