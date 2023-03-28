package config

import (
	"github.com/spf13/viper"
	"os"
)

// Server 服务器配置
var Server struct {
	HttpAddress string `mapstructure:"http_address"`
	HttpPort    int    `mapstructure:"http_port"`
}

var Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	TimeZone string `mapstructure:"timezone"`
}

var App struct {
	Mode string `mapstructure:"mode"`
}

var Auth struct {
	Key string `mapstructure:"key"`
}

var configFiles = []string{
	"./config/app.ini",
	"./config/env.ini",
}

func init() {
	for _, file := range configFiles {
		if _, err := os.Stat(file); err == nil || os.IsExist(err) {
			viper.SetConfigFile(file)
			if err := viper.MergeInConfig(); err != nil {
				panic(err)
			}
		}
	}

	if err := viper.UnmarshalKey("server", &Server); err != nil {
		panic(err)
	}

	if err := viper.UnmarshalKey("database", &Database); err != nil {
		panic(err)
	}

	if err := viper.UnmarshalKey("app", &App); err != nil {
		panic(err)
	}
}
