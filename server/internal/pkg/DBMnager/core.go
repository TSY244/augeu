package DBMnager

import (
	"augeu/server/internal/pkg/DBMnager/admin"
	"augeu/server/internal/pkg/DBMnager/task"
	"augeu/server/internal/pkg/DBMnager/volunteer"
	"augeu/server/internal/pkg/logger"
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
		&admin.Admin{},
		&volunteer.Volunteer{},
		&task.Task{},
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
