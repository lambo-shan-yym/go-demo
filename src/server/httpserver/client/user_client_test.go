package client

import "testing"

func TestGetGRClient(t *testing.T) {
	InitClient("../config/http_config_dev.ini")
	rpcClient := GetGRClient()
	if rpcClient==nil{
		t.Error("failed to get rpc client ")
	}
}
