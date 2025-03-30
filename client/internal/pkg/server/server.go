package server

import (
	config2 "augeu/client/internal/pkg/config"
	msg2 "augeu/client/internal/pkg/msg"
	_const "augeu/client/internal/utils/const"
	"augeu/public/pkg/DBMnager"
	augueMq "augeu/public/pkg/augeuMq"
	"augeu/public/pkg/logger"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"strings"
)

type Server struct {
	DbManager     *DBMnager.Manager
	RootCtx       context.Context
	Cancel        context.CancelFunc
	WebsocketConn *websocket.Conn
	Conf          *config2.Config
}

func NewServer(config *config2.Config) (*Server, error) {
	// dbm

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Db.Host, config.Db.Port, config.Db.Username, config.Db.Password, config.Db.Dbname)
	dbm, err := DBMnager.NewDBManager(dsn, true)
	if err != nil {
		log.Printf("Failed to create DB manager: %v", err)
	}
	rootCtx, cancel := context.WithCancel(context.Background())

	// mq
	if err := initMq(rootCtx); err != nil {
		logger.Errorf("Failed to init mq: %v", err)
		cancel()
		return nil, err
	}

	ws, _, err := websocket.DefaultDialer.Dial(config.Websocket.RemoteAddr, nil)
	if err != nil {
		logger.Errorf("Failed to dial websocket: %v", err)
		cancel()
		return nil, err
	}

	return &Server{
		DbManager:     dbm,
		RootCtx:       rootCtx,
		Cancel:        cancel,
		WebsocketConn: ws,
		Conf:          config,
	}, nil
}

func (s *Server) Run() {
	go s.cmdHandler()
	for {
		// 接受websocket消息
		var msg msg2.JsonMsg
		err := s.WebsocketConn.ReadJSON(&msg)
		if err != nil {
			logger.Errorf("Failed to read websocket message: %v", err)
			continue
		}
		fmt.Printf("Received message: %s\n", msg)

	}
}

// -------------------------------------- private --------------------------------------

func initMq(ctx context.Context) error {
	augueMq.Init(_const.MqMaxSize)
	err := augueMq.NewCell(ctx, _const.MqCmdName, _const.TopicMaxSize)
	if err != nil {
		return err
	}
	return nil
}

// cmd handler
func (s *Server) cmdHandler() {
	topic, err := augueMq.GetDefaultMq().GetTopic(s.RootCtx, _const.MqCmdName)
	if err != nil {
		logger.Error("msgListenerForDownloadFileSuccessQueue: ", err)
		return
	}
	for {
		select {
		case msg := <-topic.GetCell():
			cmdMsg, ok := msg.(CmdMessage)
			if !ok {
				logger.Error("msgListenerForDownloadFileSuccessQueue: msg type error")
				continue
			}
			cmd := cmdMsg.Cmd
			s.handleCmd(cmd)
		case <-s.RootCtx.Done():
			logger.Info("msgListenerForDownloadFileSuccessQueue: ctx done")
			return
		}
	}

}

func (s *Server) handleCmd(cmd string) {
	switch strings.ToLower(cmd) {
	case "exit":
		s.Cancel()
		os.Exit(0)
	default:
		logger.Error("unknown cmd: ", cmd)
	}
}
