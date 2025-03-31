package Log

import (
	"gorm.io/gorm"
	"time"
)

type EventUser struct {
	CreateAt     time.Time      `gorm:"autoCreateTime"`
	DeleteAt     gorm.DeletedAt `gorm:"index"`
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	UUID         string         `gorm:"type:varchar(255);column:uuid"`
	EventID      string         `gorm:"column:event_id"`
	SourceName   string         `gorm:"column:Source_name"`
	SourceDomain string         `gorm:"column:Source_domain"`
	TargetName   string         `gorm:"column:Target_name"`
	TargetDomain string         `gorm:"column:Target_domain"`
	MemberSid    string         `gorm:"column:MemberSid"`
	Description  string         `gorm:"type:text"`
}

func (EventUser) TableName() string {
	return "event_user"
}
