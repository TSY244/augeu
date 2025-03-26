package Log

import "time"

type EventUser struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime   time.Time `gorm:"type:timestamp"`
	EventID      string    `gorm:"column:event_id"`
	SourceName   string    `gorm:"column:Source_name"`
	SourceDomain string    `gorm:"column:Source_domain"`
	TargetName   string    `gorm:"column:Target_name"`
	TargetDomain string    `gorm:"column:Target_domain"`
	MemberSid    string    `gorm:"column:MemberSid"`
	Description  string    `gorm:"type:text"`
}

func (EventUser) TableName() string {
	return "event_user"
}
