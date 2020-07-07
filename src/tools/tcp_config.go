package tools

import (
	"fmt"
)

type TcpConfig struct {
	Db struct {
		DbDriveName       string
		DbList            string
		DbHost            string
		DbPort            string
		DbUser            string
		DbPassword        string
		DbMaxOpenConn     int
		DbMaxIdleConn     int
		DbConnMaxLifeTime int
		UseDbAccountName  string
		ShowSql           bool
	}
	Redis struct {
		Host               string
		Port               int
		Password           string
		Db                 int
		PoolSize           int
		DialTimeout        int
		IdleCheckFrequency int
		IdleTimeout        int
		MaxRetries         int
		UsedRedisServer    bool
		CacheTokenExpired  int
		CacheUserInfo      int
	}
	Server struct {
		Ip string
		Port int
	}
}

var tcpConfig *TcpConfig

func InitTcpConfig(configFile string)  {
	tcpConfig=new(TcpConfig)
	err := ConfigParser(configFile, tcpConfig)
	if err != nil {
		panic(fmt.Sprintf("failed to parser config to TcpConfig ，err:%s", err.Error()))
	}
}

func GetTcpConfig() *TcpConfig {
	return tcpConfig
}
