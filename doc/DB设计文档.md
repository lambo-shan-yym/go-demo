

### mysql设计

user_info_tab (用户信息表)

```
CREATE TABLE `user_info_tab` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL DEFAULT '',
  `nickname` varchar(64) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `profile_picture` varchar(128) NOT NULL DEFAULT '',
  `secret_key` varchar(8) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `update_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
```

字段名  |字段类型   |默认值 |说明   | 
-------- |---------|--------|----------|
id  |bigint(20) |0  |自增主键   |
username    |varchar(32)| ""    |用户名
nickname    |varchar(64)| ""    |用户昵称
password    |varchar(32)| ""    |用户密码
profile_picture |varchar(128)   | ""   |用户头像地址|
secret_key    |varchar(8)| ""    |用户秘钥
create_time |datetime   |1970-01-01 00:00:00 |创建时间
update_time |datetime   |1970-01-01 00:00:00 |更新时间|



