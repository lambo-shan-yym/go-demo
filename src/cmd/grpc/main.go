package main

import (
	"entry_task/src/server/tcpserver/api"
	"entry_task/src/tools"
)

func main() {
	// 初始化配置
	tools.InitTcpConfig("src/config/tcp_config_dev.ini")
	tcpConfig := tools.GetTcpConfig()
	// 初始化redis
	tools.InitRedis(tcpConfig)
	// 初始化数据库
	tools.InitDBConn(tcpConfig)
	// 加载日志配置
	tools.InitLog("src/config/seelog.xml")
	api.Start()
}
