package server

import (
	config2 "augeu/agent/internal/pkg/config"
	_const "augeu/agent/internal/utils/const"
	augueMq "augeu/public/pkg/augeuMq"
	"augeu/public/pkg/logger"
	"context"
	"os"
	"strings"
)

type Server struct {
	//DbManager     *DBMnager.Manager
	RootCtx context.Context
	Cancel  context.CancelFunc
	//WebsocketConn *websocket.Conn
	Conf     *config2.Config
	clientId string
}

func NewServer(config *config2.Config) (*Server, error) {
	// dbm

	//dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	config.Db.Host, config.Db.Port, config.Db.Username, config.Db.Password, config.Db.Dbname)
	//dbm, err := DBMnager.NewDBManager(dsn)
	//if err != nil {
	//	log.Printf("Failed to create DB manager: %v", err)
	//}
	rootCtx, cancel := context.WithCancel(context.Background())

	// mq
	if err := initMq(rootCtx); err != nil {
		logger.Errorf("Failed to init mq: %v", err)
		cancel()
		return nil, err
	}

	return &Server{
		//DbManager:     dbm,
		RootCtx: rootCtx,
		Cancel:  cancel,
		//WebsocketConn: ws,
		Conf: config,
	}, nil
}

func (s *Server) Run() {
	go s.cmdHandler()
	s.receiveClientId()
	select {}

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

func (s *Server) ReadMsg() {
	for {
		_, msg, err := s.WebsocketConn.ReadMessage()
		if err != nil {
			logger.Error("Failed to read message: ", err)
			continue
		}
		logger.Info("Received message: ", string(msg))
	}
}

func (s *Server) receiveClientId() {
	clientId, err := s.GetClientId()
	if err != nil {
		panic(err)
	}
	s.clientId = clientId
	logger.Info("Received client id: ", clientId)
}
