package server

import (
	"augeu/backEnd/internal/pkg/DBMnager"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
)

type Server struct {
	RootCtx context.Context
	Cancel  context.CancelFunc
	Quit    bool
	Config  *Config
	// DB manager
	DBM             *DBMnager.Manager
	WebsocketServer *websocket.Conn
	//Secrete         string // 减少查询数据库次数
}

func NewServer(config *Config) (*Server, error) {

	// dbm
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	dbm, err := DBMnager.NewDBManager(dsn)
	if err != nil {
		//log.Printf("Failed to create DB manager: %v", err)
		panic(err)
	}
	rootCtx, cancel := context.WithCancel(context.Background())
	return &Server{
		RootCtx: rootCtx,
		Cancel:  cancel,
		Config:  config,
		DBM:     dbm,
	}, nil
}
