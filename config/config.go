package config

import "os"

type AppConf struct {
	AppName string `yaml:"app_name"`
	Logger  Logger `yaml:"logger"`
}

type Logger struct {
	Level string `yaml:"level"`
}

func NewAppConf() AppConf {
	return AppConf{
		AppName: os.Getenv("APP_NAME"),
		Logger: Logger{
			Level: os.Getenv("LOG_LEVEL"),
		},
	}
}
