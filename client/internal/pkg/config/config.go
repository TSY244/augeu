package config

import "augeu/public/pkg/config"

type Config struct {
	Db        config.DbConf
	Websocket WebsocketConf
}

type WebsocketConf struct {
	RemoteAddr string
	RemotePort int
}
