package DBMnager

import (
	HostInfo2 "augeu/backEnd/internal/pkg/DBMnager/HostInfo"
	Log2 "augeu/backEnd/internal/pkg/DBMnager/Log"
	"augeu/backEnd/internal/pkg/DBMnager/NetworkInformation"
	OperateTraces2 "augeu/backEnd/internal/pkg/DBMnager/OperateTraces"
	"augeu/backEnd/internal/pkg/DBMnager/TokenTable"
	"augeu/backEnd/internal/pkg/DBMnager/UserInfo"
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

func NewDBManager(dsn string) (*Manager, error) {
	logger.Info("start to connection postgres")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &Manager{}, fmt.Errorf("DBMnage.NewDBManager -> %w ", err)
	}
	m := &Manager{
		DB: db,
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
		&HostInfo2.Patch{},
		&HostInfo2.Account{},
		&HostInfo2.System{},
		&HostInfo2.IPAddress{},
		&HostInfo2.ImageHijack{},
		&HostInfo2.ServiceInfo{},
		&HostInfo2.ScheduledTask{},
		&HostInfo2.StartupItem{},
		&HostInfo2.User{},

		// user info
		&UserInfo.UserInfo{},

		// log
		&Log2.ApplicationEvent{},
		&Log2.LoginEvent{},
		&Log2.Event7045{},
		&Log2.EventCreateProcess{},
		&Log2.EventPowerShell{},
		&Log2.EventRDPLogon{},
		&Log2.EventSIDHistory{},
		&Log2.EventUser{},
		&Log2.SecurityEvent{},
		&Log2.SystemEvent{},

		// NetworkInformation
		&NetworkInformation.ProcessConnection{},

		// OperateTraces
		&OperateTraces2.PrefetchRecord{},
		&OperateTraces2.RecentFileRecord{},
		&OperateTraces2.UserAssistRecord{},

		// ProcessInfo 暂时不添加
		//&ProcessInfo.ProcessInfo{},
		//&ProcessInfo.ProcessImport{},

		// token
		&TokenTable.Token{},
	)
	if err != nil {
		return fmt.Errorf("DBMnage.AutoMigrate -> %w", err)
	}

	// 创建索引
	if err := CreateIndex(manager.DB); err != nil {
		return fmt.Errorf("DBMnage.CreateIndex -> %w", err)
	}

	return nil
}

func CreateIndex(db *gorm.DB) error {
	err := Log2.TableUniqueConstraintsForLoginEvent(db)
	if err != nil {
		return err
	}
	err = Log2.TableUniqueConstraintsForRdpEvent(db)
	if err != nil {
		return err
	}
	err = HostInfo2.TableUniqueConstraintsForUser(db)
	if err != nil {
		return err
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
