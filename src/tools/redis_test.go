package tools

import (
	"testing"
)

func TestInitRedis(t *testing.T) {
	var tcpConfig TcpConfig
	err := ConfigParser("../../config/tcp_config_dev.ini", &tcpConfig)
	if err != nil {
		t.Error("failed to parser config to TcpConfig,err:",err.Error())
	}
	InitRedis(&tcpConfig)
	redisClient := GetRedisClient()
	if redisClient==nil{
		t.Error(" failed to init redis client")
	}
}
