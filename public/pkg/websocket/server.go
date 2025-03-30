package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Config struct {
	IsCheckOrigin    bool
	HandshakeTimeout time.Duration
	ReadBufferSize   int
	WriteBufferSize  int
}

var upgrader *websocket.Upgrader

func RegisterUpgrader(c *Config) {
	upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		HandshakeTimeout: c.HandshakeTimeout,
		ReadBufferSize:   c.ReadBufferSize,
		WriteBufferSize:  c.WriteBufferSize,
	}
}

func NewServer(listenOn string, reader func(conn *websocket.Conn), writer func(conn *websocket.Conn)) error {
	if upgrader == nil {
		log.Println("upgrader is nil")
		return nil
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		go reader(conn)
		writer(conn)
	})
	return http.ListenAndServe(listenOn, nil)
}
