package Log

import (
	"gorm.io/gorm"
	"time"
)

type Event7045 struct {
	CreateAt        time.Time      `gorm:"autoCreateTime"`
	DeleteAt        gorm.DeletedAt `gorm:"index"`
	ID              uint           `gorm:"primaryKey;autoIncrement"`
	UUID            string         `gorm:"type:varchar(255);column:uuid"`
	EventID         string         `gorm:"column:event_id"`
	ServiceName     string         `gorm:"column:Service_Name"`
	CreateAccount   string         `gorm:"column:Create_Account"`
	ServiceFilename string         `gorm:"column:Service_Filename;type:text"`
}

func (Event7045) TableName() string {
	return "event_7045"
}
