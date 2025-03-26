package Log

import "time"

type Event7045 struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	CreateTime      time.Time `gorm:"type:timestamp"`
	EventID         string    `gorm:"column:event_id"`
	ServiceName     string    `gorm:"column:Service_Name"`
	CreateAccount   string    `gorm:"column:Create_Account"`
	ServiceFilename string    `gorm:"column:Service_Filename;type:text"`
}

func (Event7045) TableName() string {
	return "event_7045"
}
