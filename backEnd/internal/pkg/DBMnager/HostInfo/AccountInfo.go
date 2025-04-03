package HostInfo

import (
	"gorm.io/gorm"
	"time"
)

// Account 账号信息表
type Account struct {
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// 使用中间表 account_ips 建立与 IPAddress 的多对多关系
	IPAddresses []IPAddress `gorm:"many2many:account_ips;"`
	ClientID    string      `gorm:"type:varchar(255);not null"` // 动态的，用于绑定任务
	// 通过 UUID 外键关联 System
	System System `gorm:"foreignKey:UUID"`
	UUID   string `gorm:"type:varchar(36);not null;primaryKey"` // UUID 通常为 36 位
}

// TableName 指定 Account 结构体对应的数据库表名
func (Account) TableName() string {
	return "accounts"
}

// IPAddress 用于存储 IP 地址
type IPAddress struct {
	gorm.Model
	Value string `gorm:"type:varchar(45);not null"` // IPv6 最长 45 位
}

// TableName 指定 IPAddress 结构体对应的数据库表名
func (IPAddress) TableName() string {
	return "ip_addresses"
}

// System 系统信息
type System struct {
	gorm.Model
	OSArch    string `gorm:"type:varchar(255);not null"`
	OSName    string `gorm:"type:varchar(255);not null"`
	OSVersion string `gorm:"type:varchar(255);not null"`
	// 使用中间表 system_patches 建立与 Patch 的多对多关系
	Patches []Patch `gorm:"many2many:system_patches;"`
	UUID    string  `gorm:"type:varchar(36);not null;uniqueIndex"` // UUID 通常为 36 位
}

// TableName 指定 System 结构体对应的数据库表名
func (System) TableName() string {
	return "systems"
}

// Patch 补丁信息
type Patch struct {
	gorm.Model
	Description string `gorm:"type:text;not null"`
	HotFixID    string `gorm:"type:varchar(255);not null;uniqueIndex"`
	InstalledBy string `gorm:"type:varchar(255);not null"`
	InstalledOn string `gorm:"type:varchar(255);not null"`
}

// TableName 指定 Patch 结构体对应的数据库表名
func (Patch) TableName() string {
	return "patches"
}

func InsertHostInfo(db *gorm.DB, account *Account) error {
	return db.Create(account).Error
}

func GetAgentInfoByClientId(db *gorm.DB, clientId string) (*Account, error) {
	var account Account
	err := db.Preload("System").Preload("IPAddresses").Where("client_id = ?", clientId).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
