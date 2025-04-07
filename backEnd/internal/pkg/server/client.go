package server

import (
	"augeu/public/pkg/logger"
	"augeu/public/pkg/websocket"
	websocket2 "github.com/gorilla/websocket"
	"sync"
	"time"
)

var (
	locker    = new(sync.RWMutex)
	ClientMap = make(map[*websocket2.Conn]string) // conn -> client id
)

func (s *Server) RunWebsocket() {
	// uint64->time.second
	timeout := time.Duration(s.Config.CoreConfig.WebsocketConf.HandshakeTimeout) * time.Second
	conf := websocket.Config{
		IsCheckOrigin:    true,
		HandshakeTimeout: timeout,
		ReadBufferSize:   s.Config.CoreConfig.WebsocketConf.ReadBufferSize,
		WriteBufferSize:  s.Config.CoreConfig.WebsocketConf.WriteBufferSize,
	}
	websocket.RegisterUpgrader(&conf)
	go websocket.NewServer(s.Config.CoreConfig.WebsocketConf.ListenOn, s.ReadMessage, s.WriteMessage)
	logger.Info("websocket server is running on ", s.Config.CoreConfig.WebsocketConf.ListenOn)
}

// 没有从client 获取数据的逻辑
func (s *Server) WriteMessage(conn *websocket2.Conn) {
	select {} // block
}

func (s *Server) ReadMessage(conn *websocket2.Conn) {
	var clientId string

	_, msg, err := conn.ReadMessage()
	if err != nil {
		logger.Errorf("read message error: %v, client id is %s", err, clientId)
		s.RemoveClient(conn)
		return
	}
	clientId = string(msg)
	logger.Info("client id: ", clientId)
	s.AddClient(conn, clientId)
	for {
		_, msg, err = conn.ReadMessage() // 只会读取一次，然后会阻塞到这里，如果对面关闭连接，就会报错，从而实现自动删除client
		if err != nil {
			logger.Errorf("read message error: %v, client id is %s", err, clientId)
			s.RemoveClient(conn)
			return
		}
		logger.Info("received message: ", string(msg))
	}
}

func (s *Server) AddClient(conn *websocket2.Conn, clientId string) {
	locker.RLock()
	for _, v := range ClientMap {
		if v == clientId {
			logger.Error("client id already exists: ", clientId)
			return
		}
	}
	locker.RUnlock()

	locker.Lock()
	defer locker.Unlock()
	ClientMap[conn] = clientId
	logger.Info("add client: ", clientId)
}

func (s *Server) RemoveClient(conn *websocket2.Conn) {
	locker.Lock()
	defer locker.Unlock()
	delete(ClientMap, conn)
}

func (s *Server) CheckClientId(clientId string) bool {
	locker.RLock()
	defer locker.RUnlock()
	for _, v := range ClientMap {
		if v == clientId {
			return true
		}
	}
	return false
}

func (s *Server) GetAllClientId() ([]string, error) {
	locker.RLock()
	defer locker.RUnlock()
	ids := make([]string, 0)
	for _, v := range ClientMap {
		ids = append(ids, v)
	}
	return ids, nil
}
