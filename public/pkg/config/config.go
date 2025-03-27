package config

type DbConf struct {
	Host       string `mapstructure:"Host"`
	Port       int    `mapstructure:"Port"`
	Username   string `mapstructure:"Username"`
	Password   string `mapstructure:"Password"`
	Dbname     string `mapstructure:"Dbname"`
	TimeOut    int    `mapstructure:"TimeOut"`
	DbFilePath string `mapstructure:"DbFilePath"`
	DbType     string `mapstructure:"DbType"`
}

type LogConf struct {
	FileName       string `mapstructure:"FileName"`
	MaxSize        int    `mapstructure:"MaxSize"`
	MaxBackups     int    `mapstructure:"MaxBackups"`
	MaxAge         int    `mapstructure:"MaxAge"`
	Compress       bool   `mapstructure:"Compress"`
	Level          string `mapstructure:"Level"`
	PrintToConsole bool   `mapstructure:"PrintToConsole"`
}
