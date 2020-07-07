package service

import (
	"entry_task/src/server/tcpserver/model"
	"entry_task/src/server/tcpserver/cache"
	"entry_task/src/server/tcpserver/dao"

	log "github.com/cihub/seelog"

	"github.com/go-redis/redis"
)

func GetUserInfo(username string) (model.User, error) {
	// 从缓存中获取用户信息
	user, err := cache.GetCacheUserInfo(username)
	if err != nil && err != redis.Nil {
		return user, err
	}
	if err == nil && user.Username == username {
		log.Debugf("【rpc服务】从缓存中获取用户信息：%v", user)
		return user, err
	}

	// 从数据库中获取用户信息
	user, err = dao.GetUserInfo(username)
	log.Debugf("【rpc服务】从数据库中获取用户信息：%v", user)
	if err != nil {
		return user, err
	}
	// 设置缓存
	err = cache.SetCacheUserInfo(user)

	if err != nil {
		log.Errorf("【rpc服务】设置用户信息到缓存中发生异常，异常信息：%s", err.Error())
	}
	return user, err
}

func UpdateUserInfo(username string, nickname string, profilePicture string, token string) int64 {
	affected, err := dao.UpdateUserInfo(username, nickname, profilePicture)
	if err != nil {
		log.Errorf("【rpc服务】更新用户信息到DB中发生异常, username:%s,nickname:%s,profilePicture:%s,err:%s", username, nickname, profilePicture, err.Error())
	}
	if affected == 1 {
		//删除缓存
		cache.DelUserInfo(username)
	}
	return affected
}
