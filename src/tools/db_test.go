package tools

import (
	"entry_task/src/server/tcpserver/model"
	"strconv"
	"testing"
)

func TestInitDBConn(t *testing.T) {

	var tcpConfig TcpConfig
	err := ConfigParser("../config/tcp_config_dev.ini", &tcpConfig)
	if err != nil {
		t.Error("failed to parser config to TcpConfig,err:", err.Error())
	}
	err = InitDBConn(&tcpConfig)
	if err != nil {
		t.Error("failed to init db, err:", err.Error())
	}
	i := 5000001
	for ; i < 10000000; i++ {
		user := &model.User{Username: strconv.Itoa(i), Password: "fec5a6bb1f7c1a8bfaa29c0839aff46c", SecretKey: "a1b2c3"}
		_, err := OrmEngine().Insert(user)
		if err != nil {
			t.Errorf("error insert %s",err.Error())
		}
	}
}
