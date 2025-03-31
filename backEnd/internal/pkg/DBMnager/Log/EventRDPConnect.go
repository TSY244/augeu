package Log

import (
	"gorm.io/gorm"
	"time"
)

type EventRDPConnect struct {
	CreateAt    time.Time      `gorm:"autoCreateTime"`
	DeleteAt    gorm.DeletedAt `gorm:"index"`
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	UUID        string         `gorm:"type:varchar(255);column:uuid"`
	EventID     string         `gorm:"column:event_id"`
	LoginName   string         `gorm:"column:LoginName"`
	Address     string         `gorm:"column:Address"`
	Domain      string         `gorm:"column:Domain"`
	Description string         `gorm:"type:text"`
}

func (EventRDPConnect) TableName() string {
	return "event_rdpconnect"
}
