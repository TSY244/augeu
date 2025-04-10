package HostInfo

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

// User 用户基础信息模型
type User struct {
	gorm.Model             // 包含ID、CreatedAt、UpdatedAt、DeletedAt字段
	Name         string    `gorm:"column:name;type:varchar(64);not null"`
	Description  string    `gorm:"column:description;type:text"`
	LocalAccount bool      `gorm:"column:local_account"`
	SID          string    `gorm:"column:sid;type:varchar(36);index"`
	IsFocus      bool      `gorm:"column:is_focus;default:false"` // 是否是可疑的用户
	UUID         uuid.UUID `gorm:"column:uuid;type:uuid"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "users"
}

// TableUniqueConstraintsForUser 创建用户表的唯一索引
func TableUniqueConstraintsForUser(db *gorm.DB) error {
	// 检查索引是否存在
	if !db.Migrator().HasIndex(&User{}, "idx_unique_user") {
		// 使用CONCURRENTLY避免锁表
		err := db.Exec(`
            CREATE UNIQUE INDEX CONCURRENTLY idx_unique_user 
            ON users 
            (uuid, sid, name )
        `).Error

		if err != nil {
			// 打印详细的错误日志
			log.Printf("创建索引失败: %v", err)
			return err
		}
	}
	return nil
}

// InsertUser 插入单条用户记录
func InsertUser(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "uuid"},
				{Name: "sid"},
				{Name: "name"},
			},
			DoNothing: true,
		}).
		Create(user).
		Error
}

// InsertUserBatch 批量插入用户记录
func InsertUserBatch(ctx context.Context, db *gorm.DB, users []User, benchSize ...int) error {
	size := 500
	if len(benchSize) > 0 {
		size = benchSize[0]
	}

	// 添加冲突处理
	return db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "uuid"},
				{Name: "sid"},
				{Name: "name"},
			},
			DoNothing: true,
		}).
		CreateInBatches(users, size).
		Error
}
