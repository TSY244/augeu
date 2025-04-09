package HostInfo

import (
	"augeu/public/pkg/logger"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

var (
	configIgnore = clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		DoNothing: true,
	}
)

// Account 账号信息表
type Account struct {
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	LastLoginTime time.Time      `gorm:"autoUpdateTime"`
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
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. 先查询是否存在记录
		var existing Account
		if err := tx.First(&existing, "uuid = ?", account.UUID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				logger.Info("查询账号记录时出错: %v", err)
				return err // 其他查询错误
			}
			// 确保关联的 Patch 记录存在
			if err := ensurePatchesExist(tx, account.System.Patches); err != nil {
				logger.Info("确保补丁记录存在时出错: %v", err)
				return err
			}
			// 手动刷新 Patch 记录的 ID
			for i, patch := range account.System.Patches {
				var existingPatch Patch
				if err := tx.First(&existingPatch, "hot_fix_id = ?", patch.HotFixID).Error; err != nil {
					logger.Info("查询补丁记录时出错: %v", err)
					return err
				}
				account.System.Patches[i].ID = existingPatch.ID
			}
			// 2. 不存在则创建（包含关联IP）
			if err := tx.Preload("System").Preload("IPAddresses").Preload("System.Patches").Create(account).Error; err != nil {
				logger.Info("创建账号记录时出错: %v", err)
				return err
			}
			return nil
		}

		// 3. 存在则执行更新（仅更新非自动字段+关联IP）
		updateFields := Account{
			ClientID:    account.ClientID,
			System:      account.System,
			IPAddresses: account.IPAddresses,
		}

		// 确保关联的 Patch 记录存在
		if err := ensurePatchesExist(tx, updateFields.System.Patches); err != nil {
			logger.Info("确保补丁记录存在时出错: %v", err)
			return err
		}
		// 手动刷新 Patch 记录的 ID
		for i, patch := range updateFields.System.Patches {
			var existingPatch Patch
			if err := tx.First(&existingPatch, "hot_fix_id = ?", patch.HotFixID).Error; err != nil {
				logger.Info("查询补丁记录时出错: %v", err)
				return err
			}
			updateFields.System.Patches[i].ID = existingPatch.ID
		}

		// 主表更新（排除自动管理字段）
		if err := tx.Model(&existing).
			Select("client_id", "system"). // 明确指定更新字段
			Updates(updateFields).Error; err != nil {
			logger.Info("更新账号记录时出错: %v", err)
			return err
		}

		// 4. 处理多对多关联（替换IP列表）
		if err := tx.Model(&existing).
			Association("IPAddresses").
			Replace(account.IPAddresses); err != nil {
			logger.Info("替换 IP 列表时出错: %v", err)
			return err
		}

		// 更新成功后同步回传最新数据
		*account = existing
		return nil
	})
}

// 确保关联的 Patch 记录存在
func ensurePatchesExist(tx *gorm.DB, patches []Patch) error {
	for _, patch := range patches {
		var existing Patch
		if err := tx.First(&existing, "hot_fix_id = ?", patch.HotFixID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 不存在则创建
				if err := tx.Create(&patch).Error; err != nil {
					logger.Info("创建补丁记录时出错: %v", err)
					return err
				}
			} else {
				logger.Info("查询补丁记录时出错: %v", err)
				return err
			}
		}
	}
	return nil
}

func GetAgentInfoByUuid(db *gorm.DB, clientId string) (*Account, error) {
	var account Account
	err := db.Preload("System").Preload("IPAddresses").Where("uuid = ?", clientId).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func GetAgentInfoByClientId(db *gorm.DB, clientId string) (*Account, error) {
	var account Account
	err := db.Preload("System").Preload("IPAddresses").Where("client_id = ?", clientId).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func GetAgentsByClientIds(db *gorm.DB, clientIds []string) ([]Account, error) {
	var accounts []Account
	err := db.Preload("System").Preload("IPAddresses").Preload("System.Patches").Where("client_id IN ?", clientIds).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
