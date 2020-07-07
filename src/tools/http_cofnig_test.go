package tools

import (
	"testing"
)

func TestGetHttpConfig(t *testing.T) {
	InitHttpConfig("http_config_dev.ini")
	config := GetHttpConfig()
	t.Logf("%v",config)
}
