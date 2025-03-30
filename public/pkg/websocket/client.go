package websocket

import "github.com/gorilla/websocket"

func Connect(server string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(server, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
