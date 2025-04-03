package Log

import (
	"augeu/public/pkg/logger"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type LoginEvent struct {
	EventId         int64          `gorm:"column:event_id"`                     // 明确指定列名
	EventTime       time.Time      `gorm:"type:timestamp(3);column:event_time"` // 时间精度到毫秒
	LoginType       string         `gorm:"type:varchar(50);column:login_type"`
	SourceIp        string         `gorm:"type:varchar(39);column:source_ip"`
	Username        string         `json:"username" gorm:"type:varchar(255)"`
	SubjectUsername string         `json:"subject_username" gorm:"type:varchar(255)"`
	SubjectDomain   string         `json:"subject_domain" gorm:"type:varchar(255)"`
	ProcessName     string         `json:"process_name" gorm:"type:varchar(255)"`
	CreateAt        time.Time      `gorm:"autoCreateTime;precision:3"` // 匹配时间精度
	DeleteAt        gorm.DeletedAt `gorm:"index"`
	ID              uint           `gorm:"autoIncrement"`
	UUID            string         `gorm:"type:varchar(255);column:uuid"`
}

func (LoginEvent) TableName() string {
	return "login_event"
}

func TableUniqueConstraints(db *gorm.DB) error {
	// 检查索引是否存在
	if !db.Migrator().HasIndex(&LoginEvent{}, "idx_unique_login_event") {
		// 使用CONCURRENTLY避免锁表
		err := db.Exec(`
            CREATE UNIQUE INDEX CONCURRENTLY idx_unique_login_event 
            ON login_event 
            (uuid, event_time, login_type, source_ip)
        `).Error

		if err != nil {
			// 打印详细的错误日志
			logger.Errorf("创建索引失败: %v", err)
			return err
		}
	}
	return nil
}

func InsertLoginEvent(ctx context.Context, db *gorm.DB, loginEvent *LoginEvent) error {
	return db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "uuid"},
				{Name: "event_time"},
				{Name: "login_type"},
				{Name: "source_ip"},
			},
			DoNothing: true,
		}).
		Create(loginEvent).
		Error
}

// InsertLoginEventBatch 保持原样

func InsertLoginEventBatch(ctx context.Context, db *gorm.DB, loginEvents []*LoginEvent, benchSize ...int) error {
	size := 500
	if len(benchSize) > 0 {
		size = benchSize[0]
	}

	// 添加冲突处理
	return db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "uuid"},
				{Name: "event_time"},
				{Name: "login_type"},
				{Name: "source_ip"},
			},
			DoNothing: true,
		}).
		CreateInBatches(loginEvents, size).
		Error
}
