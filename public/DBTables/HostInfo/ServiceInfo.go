package HostInfo

// 服务信息表
type ServiceInfo struct {
	ServiceName    string `gorm:"primaryKey;type:varchar(255)"`
	ServiceStatus  string `gorm:"type:varchar(50)"`
	Description    string `gorm:"type:text"`
	ExecutablePath string `gorm:"type:text;not null"`
	SignatureInfo  string `gorm:"type:text"`
}

func (ServiceInfo) TableName() string {
	return "service_info"
}
