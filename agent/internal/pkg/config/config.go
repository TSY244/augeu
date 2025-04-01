package config

import (
	"flag"
	"github.com/spf13/viper"
)

type Config struct {
	RemoteAddr string
	Secret     string
}

var path string
var CoreConfig Config

// 设置flag -f 指定配置文件
func init() {
	flag.StringVar(&path, "f", "./client/etc/", "config file path")
	flag.Parse()
}

func Init() (*Config, error) {
	// 使用viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&CoreConfig)
	if err != nil {
		return nil, err
	}

	return &CoreConfig, nil
}
