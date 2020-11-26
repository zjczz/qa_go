package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"likezh/api"
	"likezh/api/v1"
	"likezh/middleware"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// 中间件, 顺序不能改
	// 启动Redis的情况下将切换成Redis保存Session.
	if os.Getenv("RIM") == "use" {
		r.Use(middleware.SessionRedis(os.Getenv("SESSION_SECRET")))
	} else {
		r.Use(middleware.SessionCookie(os.Getenv("SESSION_SECRET")))
	}
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 主页
	r.GET("/", api.Index)

	// 因为v1必须依赖用户模型和Redis, 所以判断是否开启了Redis和MySQL
	if os.Getenv("RIM") != "use" {
		panic(fmt.Sprintf("v1 JWT验证必须依赖于MySQL以及Redis, 请在环境变量设置RIM为'use', 并且配置MySQL和Redis的连接"))
	}

	v1Group := r.Group("/api/v1")
	{
		// 注册
		v1Group.POST("/user/register", v1.UserRegister)
		// 登录
		v1Group.POST("/user/login", v1.UserLogin)

		// 需要登录权限
		jwt := v1Group.Group("")
		jwt.Use(middleware.JwtRequired())
		{
			// 查看个人信息
			jwt.GET("/user/me", v1.UserMe)
			// 修改密码
			jwt.POST("/user/change_password", v1.ChangePassword)
			// 退出登录
			jwt.POST("/user/logout", v1.Logout)
		}
	}

	return r
}
