package HostInfo

// 启动项表
type StartupItem struct {
	ItemName       string `gorm:"primaryKey;type:varchar(255)"`
	ExecutablePath string `gorm:"type:text;not null"`
	SignatureInfo  string `gorm:"type:text"`
}

func (StartupItem) TableName() string {
	return "startup_items"
}
