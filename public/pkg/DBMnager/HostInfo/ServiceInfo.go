package HostInfo

import (
	"gorm.io/gorm"
	"time"
)

// 服务信息表
type ServiceInfo struct {
	ServiceName    string         `gorm:"primaryKey;type:varchar(255)"`
	ServiceStatus  string         `gorm:"type:varchar(50)"`
	Description    string         `gorm:"type:text"`
	ExecutablePath string         `gorm:"type:text;not null"`
	SignatureInfo  string         `gorm:"type:text"`
	CreateAt       time.Time      `gorm:"autoCreateTime"`
	DeleteAt       gorm.DeletedAt `gorm:"index"`
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	UUID           string         `gorm:"type:varchar(255);column:uuid"`
}

func (ServiceInfo) TableName() string {
	return "service_info"
}
