package cache

import (
	"encoding/json"
	"entry_task/src/server/tcpserver/model"
	"entry_task/src/tools"
	"time"
)

const (
	userInfoPrefix  = "entry_task_user_info_"
	userTokenPrefix = "entry_task_user_token_"
)

// 根据用户名从缓存中获取用户信息
func GetCacheUserInfo(username string) (model.User, error) {
	redisKey := userInfoPrefix + username
	val, err := tools.GetRedisClient().Get(redisKey).Result()
	var user model.User
	if err != nil {
		return user, err
	}
	err = json.Unmarshal([]byte(val), &user)
	return user, err
}

// 设置用户信息到缓存中
func SetCacheUserInfo(user model.User) error {
	redisKey := userInfoPrefix + user.Username

	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	// 根据配置设置过期时间
	expire := time.Second * time.Second * time.Duration(tools.GetTcpConfig().Redis.CacheUserInfo)
	_, err = tools.GetRedisClient().Set(redisKey, bytes, expire).Result()
	return err
}

//

func UpdateCacheUserInfo(user model.User) error {
	err := SetCacheUserInfo(user)
	if err != nil {
		redisKey := userInfoPrefix + user.Username
		tools.GetRedisClient().Del(redisKey)
	}
	return err
}

func DelUserInfo(username string)  error{
	redisKey := userInfoPrefix + username
	_, err := tools.GetRedisClient().Del(redisKey).Result()
	return err
}