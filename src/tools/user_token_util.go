package tools

import (
	"fmt"
	"time"
	"usermm/utils"
)

func GenerateToken(username string) string {
	// 获取当前时间的时间戳
	t := time.Now().Unix()
	return utils.Md5String(fmt.Sprintf("%s:%d",username,t))
}


func GenerateTokenForTest(username string) string {
	// 获取当前时间的时间戳
	return GetMd5String(username)
}

