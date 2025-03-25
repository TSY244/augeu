package server

import (
	"augeu/server/internal/pkg/DBMnager"
	"context"
	"fmt"
	"log"
)

type Server struct {
	RootCtx context.Context
	Cancel  context.CancelFunc
	Quit    bool
	Config  *Config
	// DB manager
	DBM *DBMnager.Manager
}

func NewServer(config *Config) (*Server, error) {
	// dbm
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	dbm, err := DBMnager.NewDBManager(dsn)
	if err != nil {
		log.Printf("Failed to create DB manager: %v", err)
	}
	rootCtx, cancel := context.WithCancel(context.Background())
	return &Server{
		RootCtx: rootCtx,
		Cancel:  cancel,
		Config:  config,
		DBM:     dbm,
	}, nil
}
