package config

import "augeu/public/pkg/config"

var (
	CoreConfig = Config{}
)

type Config struct {
	HttpServer    HttpServerConf `mapstructure:"HttpServerConf"`
	DbConfig      config.DbConf  `mapstructure:"DatabaseConf"`
	LogConfig     config.LogConf `mapstructure:"LogConf"`
	WebsocketConf WebsocketConf  `mapstructure:"WebsocketConf"`
}

type HttpServerConf struct {
	HttpPort         int    `mapstructure:"Port"`
	HttpListen       string `mapstructure:"ListenOn"`
	EnableStaticAuth bool   `mapstructure:"EnableStaticAuth"`
}

type WebsocketConf struct {
	ListenOn         string `mapstructure:"ListenOn"`
	MaxSize          int    `mapstructure:"MaxSize"`
	ReadBufferSize   int    `mapstructure:"ReadBufferSize"`
	WriteBufferSize  int    `mapstructure:"WriteBufferSize"`
	WriteWait        int    `mapstructure:"WriteWait"`
	PongWait         int    `mapstructure:"PongWait"`
	PingPeriod       int    `mapstructure:"PingPeriod"`
	MaxMessageSize   int    `mapstructure:"MaxMessageSize"`
	HandshakeTimeout uint64 `mapstructure:"HandshakeTimeout"`
}
