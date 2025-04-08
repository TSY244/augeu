package Log

import (
	"augeu/public/pkg/logger"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type EventRDPLogon struct {
	CreateAt      time.Time      `gorm:"autoCreateTime"`
	DeleteAt      gorm.DeletedAt `gorm:"index"`
	ID            uint           `gorm:"primaryKey;autoIncrement"`
	UUID          string         `gorm:"type:varchar(255);column:uuid"`
	EventID       int64          `gorm:"column:event_id"`
	AccountDomain string         `gorm:"column:account_domain"`
	AccountName   string         `gorm:"column:account_name"`
	ClientAddress string         `gorm:"column:client_address"`
	ClientName    string         `gorm:"column:client_name"`
}

func (EventRDPLogon) TableName() string {
	return "event_rdplogon"
}

func TableUniqueConstraintsForRdpEvent(db *gorm.DB) error {
	// 检查索引是否存在
	if !db.Migrator().HasIndex(&EventRDPLogon{}, "idx_unique_rdp_event") {
		// 使用CONCURRENTLY避免锁表
		err := db.Exec(`
            CREATE UNIQUE INDEX CONCURRENTLY idx_unique_rdp_event 
            ON event_rdplogon 
            (uuid, account_domain, account_name, client_address,client_name)
        `).Error

		if err != nil {
			// 打印详细的错误日志
			logger.Errorf("创建索引失败: %v", err)
			return err
		}
	}
	return nil
}

func InsertRdpEvent(ctx context.Context, db *gorm.DB, rdpEvent *EventRDPLogon) error {
	return db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "uuid"},
				{Name: "account_domain"},
				{Name: "account_name"},
				{Name: "client_address"},
				{Name: "client_name"},
			},
			DoNothing: true,
		}).
		Create(rdpEvent).
		Error
}

func InsertRdpEventBatch(ctx context.Context, db *gorm.DB, rdpEvents []EventRDPLogon, benchSize ...int) error {
	size := 500
	if len(benchSize) > 0 {
		size = benchSize[0]
	}

	// 添加冲突处理
	return db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "uuid"},
				{Name: "account_domain"},
				{Name: "account_name"},
				{Name: "client_address"},
				{Name: "client_name"},
			},
			DoNothing: true,
		}).
		CreateInBatches(rdpEvents, size).
		Error
}
