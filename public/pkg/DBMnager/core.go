package DBMnager

import (
	"augeu/public/pkg/DBMnager/HostInfo"
	"augeu/public/pkg/DBMnager/Log"
	"augeu/public/pkg/DBMnager/NetworkInformation"
	"augeu/public/pkg/DBMnager/OperateTraces"
	"augeu/public/pkg/logger"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Manager struct {
	DB *gorm.DB
}

func NewDBManager(dsn string, isClient bool) (*Manager, error) {
	logger.Info("start to connection postgres")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &Manager{}, fmt.Errorf("DBMnage.NewDBManager -> %w ", err)
	}
	m := &Manager{
		DB: db,
	}
	if isClient {
		return m, nil
	}
	err = m.AutoMigrate()
	if err != nil {
		return m, fmt.Errorf("DBMnage.NewDBManager -> %w", err)
	}

	return m, nil
}

func (manager *Manager) AutoMigrate() error {
	logger.Info("starting to migrate databases...")
	err := manager.DB.AutoMigrate(
		// host info
		&HostInfo.AccountInfo{},
		&HostInfo.ImageHijack{},
		&HostInfo.ServiceInfo{},
		&HostInfo.ScheduledTask{},
		&HostInfo.StartupItem{},

		// log
		&Log.ApplicationEvent{},
		&Log.LoginEvent{},
		&Log.Event7045{},
		&Log.EventCreateProcess{},
		&Log.EventPowerShell{},
		&Log.EventRDPConnect{},
		&Log.EventRDPLogon{},
		&Log.EventSIDHistory{},
		&Log.EventUser{},
		&Log.SecurityEvent{},
		&Log.SystemEvent{},

		// NetworkInformation
		&NetworkInformation.ProcessConnection{},

		// OperateTraces
		&OperateTraces.PrefetchRecord{},
		&OperateTraces.RecentFileRecord{},
		&OperateTraces.UserAssistRecord{},

		// ProcessInfo 暂时不添加
		//&ProcessInfo.ProcessInfo{},
		//&ProcessInfo.ProcessImport{},
	)
	if err != nil {
		return fmt.Errorf("DBMnage.AutoMigrate -> %w", err)
	}
	return nil
}

// isUniqueConstraintError checks if the error is a unique constraint error
func isUniqueConstraintError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}
