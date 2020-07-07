package tools

import (
	"fmt"
)

type HttpConfig struct {
	Server struct {
		Port int
	}
	Rpc struct {
		Address string
	}
	Image struct {
		PrefixUrl string
		SavePath  string
	}
}

var httpConfig *HttpConfig

func InitHttpConfig(configFile string) {
	httpConfig = new(HttpConfig)
	err := ConfigParser(configFile, httpConfig)
	if err != nil {
		panic(fmt.Sprintf("failed to parser config to TcpConfig ，err:%s", err.Error()))
	}
}

func GetHttpConfig() *HttpConfig {
	return httpConfig
}
