package model

type User struct {
	Id             int32
	Username       string `xorm:"varchar(32) not null unique" `
	Nickname       string `xorm:"varchar(64)"`
	Password       string `xorm:"varchar(32)"`
	ProfilePicture string `xorm:"varchar(256)"`
	SecretKey      string `xorm:varchar(8)`
	CreateTime   uint32 `xorm:"int(11) unsigned not null default '0' 'create_time'"`
	UpdateTime   uint32 `xorm:"int(11) unsigned not null default '0' 'update_time'"`
}
