package server

import (
	config2 "augeu/agent/internal/pkg/config"
	"augeu/agent/internal/pkg/msg"
	"augeu/agent/internal/pkg/systeminfo"
	_const "augeu/agent/internal/utils/const"
	"augeu/agent/internal/utils/utils"
	augueMq "augeu/public/pkg/augeuMq"
	"augeu/public/pkg/logger"
	utils2 "augeu/public/util/utils"
	"context"
	"github.com/gorilla/websocket"
	"os"
	"strings"
)

type Server struct {
	//DbManager     *DBMnager.Manager
	RootCtx       context.Context
	Cancel        context.CancelFunc
	WebsocketConn *websocket.Conn
	Conf          *config2.Config
	clientId      string
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

	ws, _, err := websocket.DefaultDialer.Dial(config.Websocket.RemoteAddr, nil)
	if err != nil {
		logger.Errorf("Failed to dial websocket: %v", err)
		cancel()
		return nil, err
	}

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

	structmsg, err := systeminfo.GetSystemInfo()
	if err != nil {
		panic(err)
	}
	uuid, err := systeminfo.GetUuid()
	if err != nil {
		panic(err)
	}
	ips, err := utils2.GetIps()
	if err != nil {
		panic(err)
	}
	tempData := msg.HelloMsg{
		UUID:       uuid,
		IP:         ips,
		SystemInfo: *structmsg,
	}
	msgStr, err := utils.StructToJson(tempData)
	if err != nil {
		panic(err)
	}
	baseMsg := msg.JsonMsg{
		Type:     msg.MessageType,
		ClientId: s.clientId,
		Message:  msgStr,
	}
	baseMsgStr, err := utils.StructToJson(baseMsg)
	if err != nil {
		panic(err)
	}
	err = s.WebsocketConn.WriteMessage(websocket.TextMessage, []byte(baseMsgStr))
	if err != nil {
		panic(err)
	}

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

}
