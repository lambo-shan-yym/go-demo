package dao

import (
	"entry_task/src/server/tcpserver/model"
	"entry_task/src/tools"
)

func GetUserInfo(username string) (model.User, error) {
	user := model.User{}
	_, err := tools.OrmEngine().Where("username=?", username).Get(&user)
	if err != nil {
		return user, err
	}
	return user, nil;
}

func UpdateUserInfo(username, nickname, profilePicture string) (int64, error) {
	user := &model.User{}
	if nickname != "" {
		user.Nickname = nickname
	}
	if profilePicture != "" {
		user.ProfilePicture = profilePicture
	}
	return tools.OrmEngine().Where("username=?", username).Update(user)

}

func InsertUser(username string, password string) (int64, error) {
	user := &model.User{Username: username, Password: password}
	return tools.OrmEngine().Insert(user)
}
