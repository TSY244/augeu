package HostInfo

// 账号信息表
type AccountInfo struct {
	AccountName  string `gorm:"primaryKey;type:varchar(255);column:account_name"`
	IsSuspicious bool   `gorm:"not null;default:false"`
}

func (AccountInfo) TableName() string {
	return "account_info"
}
