package server

import (
	config2 "augeu/agent/internal/pkg/config"
	_const "augeu/agent/internal/utils/const"
	"augeu/agent/pkg/windowsLog"
	augueMq "augeu/public/pkg/augeuMq"
	"augeu/public/pkg/logger"
	"context"
	"github.com/gorilla/websocket"
	"os"
	"strings"
	"time"
)

type Server struct {
	//DbManager     *DBMnager.Manager
	RootCtx       context.Context
	Cancel        context.CancelFunc
	WebsocketConn *websocket.Conn
	Conf          *config2.Config
	ClientId      string
	Jwt           string
	Header        map[string]string
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

	// websocket
	ws, resp, err := websocket.DefaultDialer.Dial(config.WebsocketAddr, nil)
	if err != nil {
		logger.Errorf("Failed to dial websocket: %v", err)
		cancel()
		return nil, err
	}
	if resp.StatusCode != 101 {
		logger.Errorf("Failed to dial websocket: %v", resp.Status)
		cancel()
		return nil, err
	}
	go func() {
		for {
			select {
			case <-rootCtx.Done():
				logger.Info("websocket connection closed")
				return
			default:
			}

			_, _, err := ws.ReadMessage()
			if err != nil {
				logger.Errorf("Lost Connection to server: %v", err)
				cancel()
				return
			}

		}
	}()

	return &Server{
		//DbManager:     dbm,
		RootCtx:       rootCtx,
		Cancel:        cancel,
		WebsocketConn: ws,
		Conf:          config,
	}, nil
}

func (s *Server) Run() {
	go s.cmdHandler()
	s.receiveClientId() // 获取 ClientId
	s.sendClientId()    // 建立websocket连接
	s.core()

}

// -------------------------------------- private --------------------------------------

func (s *Server) core() {
	windowsLog.RegisterFunctionMap(windowsLog.LoginEvenType, s.PushLoginEvent)
	//for {
	err := s.GetLoginEventInfo()
	if err != nil {
		logger.Errorf("Failed to get login event info: %v", err)
		time.Sleep(10 * time.Second)
		//continue
	}

	time.Sleep(10 * time.Second)
	//}
}

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

func (s *Server) sendClientId() {
	clientId := s.ClientId
	if clientId == "" {
		logger.Error("ClientId is empty")
		return
	}
	err := s.WebsocketConn.WriteMessage(websocket.TextMessage, []byte(clientId))
	if err != nil {
		logger.Error("Failed to send client id: ", err)
		return
	}
	logger.Info("Sent client id: ", clientId)
}

func (s *Server) receiveClientId() {
	jwt, clientId, err := s.GetClientId()
	if err != nil {
		panic(err)
	}
	s.ClientId = clientId
	s.Jwt = jwt
	s.Header = map[string]string{
		"Authorization": jwt,
	}
	logger.Info("Received client id: ", clientId)
	logger.Info("Received jwt: ", jwt)
}
