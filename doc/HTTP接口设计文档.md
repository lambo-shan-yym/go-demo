# entry task http接口设计计
版本:
 
版本 | 日期      |操作人
--- |---|---
v1  | 2020-05-09|杨阳明



HTTP API

----
## 1、用户登录接口
###  【接口url】
login
###  【请求方式】
POST

### 【请求头】
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|
###  【请求body(JSON)】
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|
username    | string  |是       |用户名    | 
password    | string  |是       |密码(明文密码md5加密过后)|

### 【请求示例】
```
{
    "username": "lambo",
    "password" : "xxxxx"
   
}
```
### 【接口响应】
 
返回结果参数中字段说明：
 
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""|
| result | json | N    |
          
result

 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
token    | string | Y      | 用户token|

### 【错误码】
 code      |msg      |desc       | 
-------- |---------|--------|
0    | "" | success      |


### 【返回结果示例】
```
{
    "coode": 0,
    "msg": "",
    "result": {
        "nickname": "亖狼何必装羴2",
        "profilePicture": "http://localhost:8080/front_end/upload/38c32e17995ed5465e2a0221418dd30f.jpg",
        "token": "6ab3f8e9ef36cdf3036673ee2934eb9a",
        "username": "lambo"
    }
}
```

----
## 2、用户退出登录接口
###  【接口url】
user/logout
###  【请求方式】
POST

### 【请求头】
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|
Authorization    | string  |Y       |用户token    |

### 【请求body】
name      |type     |required |desc       | 
-------- |---------|--------|----------|


### 【请求示例】
header内容

```
{
    "Authorization":"6ab3f8e9ef36cdf3036673ee2934eb9a"
}
```
### 【接口响应】
 
返回结果参数中字段说明：
 
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""|
| result | json | N    |
          
result

 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|

### 【错误码】
 code      |msg      |desc       | 
-------- |---------|--------|
0    | "" | success      |


### 【返回结果示例】
```
{
    "coode": 0,
    "msg": "",
    "result":""
}
```



----
## 3、修改用户信息接口
###  【接口url】
user/update_user_info
###  【请求方式】
PUT

### 【请求头】
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|
Authorization    | string  |Y       |用户token    |


### 【请求body】
name      |type     |required |desc       | 
-------- |---------|--------|----------|
nickname |string |Y|用户昵称

### 【请求示例】
header内容
```
{
    "Authorization":"6ab3f8e9ef36cdf3036673ee2934eb9a"
}
```
body参数
```
{
	"nickname":"test"
}
```
### 【接口响应】
 
返回结果参数中字段说明：
 
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""|
| result | json | N    |
          
result

 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|

### 【错误码】
 code      |msg      |desc       | 
-------- |---------|--------|
0    | "" | success      |


### 【返回结果示例】
```
{
    "coode": 0,
    "msg": "",
    "result":""
}
```

----
## 4、修改用户头像接口
###  【接口url】
user/update_user_profile_picture
###  【请求方式】
PUT

### 【请求头】
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|
Authorization    | string  |Y       |用户token    |

### 【url参数】
name      |type     |required |desc       | 
-------- |---------|--------|----------|

### 【请求表单】
name      |type     |required |desc       | 
-------- |---------|--------|----------|
file | file | Y |用户头像

### 【请求示例】
header 内容
```
{
    "Authorization":"6ab3f8e9ef36cdf3036673ee2934eb9a"
}
```

form表单
```
"file":图片
```
### 【接口响应】
 
返回结果参数中字段说明：
 
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""|
| result | json | N    |
          
result

 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|

### 【错误码】
 code      |msg      |desc       | 
-------- |---------|--------|
0    | "" | success      |


### 【返回结果示例】
```
{
    "coode": 0,
    "msg": "",
    "result":""
}
```


----
## 5、获取用户信息接口
###  【接口url】
user/get_user_info
###  【请求方式】
GET

### 【请求头】
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|
Authorization    | string  |Y       |用户token    |

### 【url参数】
name      |type     |required |desc       | 
-------- |---------|--------|----------|

### 【请求body】
name      |type     |required |desc       | 
-------- |---------|--------|----------|


### 【请求示例】

header内容
```
{
    "Authorization":"6ab3f8e9ef36cdf3036673ee2934eb9a"
}
```
### 【接口响应】
 
返回结果参数中字段说明：
 
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""|
| result | json | N    |
          
result

 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
username | string | Y|用户名
nickname | string | Y| 用户昵称
profile_picture|string | Y|用户头像

### 【错误码】
 code      |msg      |desc       | 
-------- |---------|--------|
0    | "" | success      |


### 【返回结果示例】
```
{
    "code": 0,
    "msg": "",
    "result": {
        "nickname": "test",
        "profile_picture": "http://localhost:8080/src/front_end/upload/e830be218f9c4f983cc001d52c893b99.jpg",
        "username": "10340"
    }
}
```


