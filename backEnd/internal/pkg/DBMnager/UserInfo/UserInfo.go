package UserInfo

import (
	"augeu/public/pkg/encoding"
	"errors"
	"gorm.io/gorm"
)

// UserInfo 用户信息表
type UserInfo struct {
	gorm.Model        // 包含ID、CreatedAt、UpdatedAt、DeletedAt字段
	UserName   string `gorm:"column:user_name;type:varchar(64);primarykey;not null" json:"user_name"` // 用户名（唯一索引）
	Password   string `gorm:"column:password;type:varchar(255);not null" json:"-"`                    // 密码（存储哈希值）
}

// TableName 自定义表名
func (UserInfo) TableName() string {
	return "user_infos" // 表名使用复数形式
}

func AddUser(userName, password string, db *gorm.DB) error {
	return db.Create(&UserInfo{
		UserName: userName,
		Password: encoding.Md5Hash(password),
	}).Error
}

func CheckUser(db *gorm.DB, userName, password string) error {
	var user UserInfo
	err := db.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return err
	}
	if user.Password != encoding.Md5Hash(password) {
		return errors.New("用户名或密码错误")
	}
	return nil
}
