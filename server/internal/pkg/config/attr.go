package config

import "augeu/public/pkg/config"

var (
	CoreConfig = Config{}
)

type Config struct {
	HttpServer HttpServerConf `mapstructure:"HttpServerConf"`
	DbConfig   config.DbConf  `mapstructure:"DatabaseConf"`
	LogConfig  config.LogConf `mapstructure:"LogConf"`
}

type HttpServerConf struct {
	HttpPort         int    `mapstructure:"Port"`
	HttpListen       string `mapstructure:"ListenOn"`
	EnableStaticAuth bool   `mapstructure:"EnableStaticAuth"`
}
