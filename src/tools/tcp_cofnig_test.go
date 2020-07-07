package tools

import (
	"testing"
)

func TestGetTcpConfig(t *testing.T) {
	InitTcpConfig("tcp_config_dev.ini")
	config := GetTcpConfig()
	t.Logf("%v",config)
}
