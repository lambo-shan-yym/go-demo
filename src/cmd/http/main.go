package main

import (
	"entry_task/src/server/httpserver/client"
	"entry_task/src/server/httpserver/router"
	"entry_task/src/tools"
)

func main() {
	client.InitClient("src/config/http_config_dev.ini")
	tools.InitLog("src/config/seelog.xml")
	r := router.InitRouter()
	panic(r.Run())
}
