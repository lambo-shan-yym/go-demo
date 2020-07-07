
# entry task rpc接口设计
版本:
 
版本 | 日期      |操作人
--- |---|---
v1  | 2020-05-09|杨阳明


TCP API
----
## 1、用户登录接口
【api_request_function】 :  Login

【api_request_description】 : 

【api_request_protocol】 : TCP

【api_request_method】 :  gRPC + protobuf3
【api_request_params】:
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|
username    | string  |Y       |用户名    |       | 
password    | string  |Y       |密码(明文密码md5加密过后)|

【api_response_params】:
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""|
          


## 2、用户退出登录接口
【api_request_function】 :  Logout

【api_request_description】 : 

【api_request_protocol】 : TCP

【api_request_method】 :  gRPC + protobuf3
【api_request_params】:
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|

【api_response_params】:
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""|






## 3、获取用户信息接口
【api_request_function】 :  GetUserInfo

【api_request_description】 : 

【api_request_protocol】 : TCP

【api_request_method】 :  gRPC + protobuf3
【api_request_params】:
 
name      |type     |required |desc       | 
-------- |---------|--------|----------|


【api_response_params】:
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""
username | string | N|用户名
nickname | string | N| 用户昵称
profile_picture|string | N|用户头像



## 4、修改个人信息接口
【api_request_function】 :  EditUserInfo

【api_request_description】 : 

【api_request_protocol】 : TCP

【api_request_method】 :  gRPC + protobuf3
【api_request_params】:
 
name      |type     |required |desc       | 
-------- |---------|--------|----------| 
nickname    | string  |N       |用户昵称
profile_picture    | string  |N       |个人头像地址|
【api_response_params】:
 
 name      |type     |required |desc       | 
-------- |---------|--------|----------|
code    | unit32  |Y      |error code: success = 0 , fail > 0    |
| msg    | string  |N       |error message: success = ""