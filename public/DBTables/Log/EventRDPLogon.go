package Log

import "time"

type EventRDPLogon struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime  time.Time `gorm:"type:timestamp"`
	EventID     string    `gorm:"column:event_id"`
	LoginName   string    `gorm:"column:LoginName"`
	Address     string    `gorm:"column:Address"`
	Domain      string    `gorm:"column:Domain"`
	Description string    `gorm:"type:text"`
}

func (EventRDPLogon) TableName() string {
	return "event_rdplogon"
}
