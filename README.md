## 项目目标
* 让团队更好地了解新人对技能的掌握情况
* 熟悉使用Go gin框架
* 熟悉使用Go gRPC框架
* 熟悉基于Auth Token的鉴权机制和流程
* 熟悉使用Go对MySQL、Redis进行基本操作
* 合理安排任务进度和时间
* 熟悉代码规范、测试、文档、性能调优

## 项目功能需求
* 现一个用户管理系统，用户可以登录、拉取和编辑他们的profiles。
* 用户可以通过在Web页面输入username和password登录，backend系统负责校验用户身份。成功登录后，页面需要展示用户的相关信息；否则页面展示相关错误。
* 成功登录后，用户可以编辑以下内容：
   1. 上传profile picture
   2. 修改nickname（需要支持Unicode字符集，utf-8编码）
  
* 用户信息包括：
  1. username（不可更改）
  2. nickname
  3. profile picture
* 需要提前将初始用户数据插入数据库用于测试。确保测试数据库中包含10,000,000条用户账号信息。
## 项目目录结构
```
├── doc                             // 文档
├── log                             // 日志输出目录
├── src                             // 源码目录
│   ├── cmd                         // main函数启动入口
│   │   ├── grpc                    // grpc服务main函数启动入口
│   │   ├── http                    // http服务main函数入口
│   ├── config                      // 配置文件
│   ├── front_end                   // 前端代码
│   ├── proto                       // 
│   ├── server                      // 服务
│   │   ├── httpserver              // http服务
│   │   │   ├── api                 // 接口
│   │   │   ├── client              // 客户端
│   │   │   ├── code                // 异常码
│   │   │   ├── dto                 // 数据传输对象
│   │   │   ├── router              // 路由
│   │   ├── tcpserver               // tcp服务
│   │   │   ├── api                 // grpc接口
│   │   │   ├── cache               // 缓存
│   │   │   ├── code                // 异常码
│   │   │   ├── dao                 // dao层
│   │   │   ├── model               // 实体
│   │   │   ├── service             // 业务逻辑
│   ├── tools                       // 工具类
```



## 相关文档

[DB设计文档](https://git.garena.com/yangming.yang/entry_task/blob/dev/doc/DB%E8%AE%BE%E8%AE%A1%E6%96%87%E6%A1%A3.md)

[RPC接口设计文档](https://git.garena.com/yangming.yang/entry_task/blob/dev/doc/RPC%E6%8E%A5%E5%8F%A3%E8%AE%BE%E8%AE%A1%E6%96%87%E6%A1%A3.md)

[HTTP接口设计文档](https://git.garena.com/yangming.yang/entry_task/blob/dev/doc/HTTP%E6%8E%A5%E5%8F%A3%E8%AE%BE%E8%AE%A1%E6%96%87%E6%A1%A3.md)

[压测结果文档](https://git.garena.com/yangming.yang/entry_task/blob/dev/doc/%E6%80%A7%E8%83%BD%E6%B5%8B%E8%AF%95%E7%BB%93%E6%9E%9C.md)

[总结文档](https://git.garena.com/yangming.yang/entry_task/blob/dev/doc/%E6%80%BB%E7%BB%93%E6%96%87%E6%A1%A3.md)



## 前端效果
### 登录界面  
![image](https://github.com/lambo-shan-yym/go-entry-task/blob/master/doc/%E5%9B%BE%E7%89%87/login.png)

### 登录失败提示
![image](https://github.com/lambo-shan-yym/go-entry-task/blob/master/doc/%E5%9B%BE%E7%89%87/login_error.png)

### 个人信息页面
![image](https://github.com/lambo-shan-yym/go-entry-task/blob/master/doc/%E5%9B%BE%E7%89%87/user_info.png)
### 修改个人基本信息成功提示
![image](https://github.com/lambo-shan-yym/go-entry-task/blob/master/doc/%E5%9B%BE%E7%89%87/update_user_info_success.png)
### 修改个人头像成功提示
![image](https://github.com/lambo-shan-yym/go-entry-task/blob/master/doc/%E5%9B%BE%E7%89%87/update_user_profile_picture_success.png)