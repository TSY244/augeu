package server

import (
	config2 "augeu/client/internal/pkg/config"
	"augeu/public/pkg/DBMnager"
	"context"
	"fmt"

	"log"
)

type Server struct {
	DbManager *DBMnager.Manager
	RootCtx   context.Context
	Cancel    context.CancelFunc
}

func NewServer(config *config2.Config) (*Server, error) {
	// dbm
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Db.Host, config.Db.Port, config.Db.Username, config.Db.Password, config.Db.Dbname)
	dbm, err := DBMnager.NewDBManager(dsn, true)
	if err != nil {
		log.Printf("Failed to create DB manager: %v", err)
	}
	rootCtx, cancel := context.WithCancel(context.Background())
	return &Server{
		DbManager: dbm,
		RootCtx:   rootCtx,
		Cancel:    cancel,
	}, nil
}
