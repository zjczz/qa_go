package routes

import (
	"likezh/api"
	v1 "likezh/api/v1"
	"likezh/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	// 主页
	r.GET("/", api.Index)

	v1Group := r.Group("/api/v1")
	{
		// 注册
		v1Group.POST("/user/register", v1.UserRegister)
		// 登录
		v1Group.POST("/user/login", v1.UserLogin)

		// 查看单个问题
		v1Group.GET("/questions/:id", v1.FindOneQuestion)

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
			// 发布问题
			jwt.POST("/questions", v1.QuestionAdd)
		}
	}

	return r
}
