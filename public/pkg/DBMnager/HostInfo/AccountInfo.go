package HostInfo

import "gorm.io/gorm"

// 账号信息表
type AccountInfo struct {
	gorm.Model
	AccountName  string `gorm:"type:varchar(255);column:account_name"`
	IsSuspicious bool   `gorm:"not null;default:false"`
	UUID         string `gorm:"primaryKey;type:varchar(255);column:uuid"`
}

func (AccountInfo) TableName() string {
	return "account_info"
}
