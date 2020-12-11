# qa_go

工程实践，一个类似知乎的问答系统（服务端），正在开发中。。

- [x] 用户注册
- [x] 用户登录
- [x] 查看个人信息
- [x] 发布问题
- [x] 查看问题列表
- [x] 查看单个问题信息
- [ ] 修改问题
- [ ] 删除问题
- [ ] 回答问题
- [ ] 修改回答
- [ ] 删除回答
- [ ] 热门问题列表
- [ ] 回答点赞点踩
- [ ] 查看自己发布问题，回答问题，评论记录

## 环境依赖

- [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架
- [GORM](http://gorm.io/docs/index.html): ORM工具，本项目需要配合MySQL使用
- [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端，用于缓存相关功能
- [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便配置环境变量
- [Jwt-Go](https://github.com/dgrijalva/jwt-go): Golang JWT组件，本项目使用基于 jwt 实现的 token 来做身份验证
- [crypto](https://pkg.go.dev/golang.org/x/crypto): Golang 加密算法库，本项目使用其中的 bcrypto 算法来加密用户密码

## 目录结构

```
├── api              controller层，负责处理请求
│   ├── v1           api版本
├── cache            主要是 redis 缓存相关 
├── conf             项目的静态配置
├── middleware       中间件
├── model            数据库模型以及相关操作
├── routes           路由配置
├── serializer       将实体映射成不同的viewmodel，以及常用的响应信息
├── service          将比较复杂的业务从api层分离出来
├── utils            常用工具
```

## 使用方法

### 1. 获取代码

```
git clone https://github.com/zjczz/qa_go
```

### 2. 修改环境变量配置

将项目目录下的 **example.env** 复制一份，命名为 **.env** ，修改其中 MySQL，Redis 相关配置，JWT_SECRET_KEY 是用于 jwt 加密的秘钥，设置为一个随机字符串即可

### 3. 直接运行

确保 golang 版本在 1.14 及以上并开启了 **go module** ，执行 `go run main.go` 便会自动下载依赖并启动项目

### 4. docker部署

如果你安装了 docker，可以很方便的用 docker 部署本项目，见 Dockerfile 和 docker-compose.yml
