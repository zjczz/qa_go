package routes

import (
	"qa_go/api"
	v1 "qa_go/api/v1"

	"qa_go/middleware/auth"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 主页
	r.GET("/", api.Index)

	v1Group := r.Group("/api/v1")
	v1Group.Use(auth.JwtWithAnonymous())
	{
		// 注册
		v1Group.POST("/user/register", v1.UserRegister)
		// 登录
		v1Group.POST("/user/login", v1.UserLogin)

		// 查看单个问题
		v1Group.GET("/questions/:qid", v1.FindOneQuestion)

		// 获取热门问题列表
		v1Group.GET("/hot_questions", v1.FindHotQuestions)

		// 获取首页问题列表
		v1Group.GET("/questions", v1.FindQuestions)

		// 查看单个回答
		v1Group.GET("/questions/:qid/answers/:aid", v1.FindAnswer)
		

		// 需要登录权限
		jwt := v1Group.Group("")
		jwt.Use(auth.JwtRequired())
		{
			// 查看个人信息
			jwt.GET("/user/me", v1.UserMe)

			// 发布问题
			jwt.POST("/questions", v1.QuestionAdd)
			// 修改问题
			jwt.PUT("/questions/:qid", v1.EditQuestion)
			// 删除问题
			jwt.DELETE("/questions/:qid", v1.DeleteQuestion)

			// 回答问题
			jwt.POST("/questions/:qid/answers", v1.AddAnswer)
		}
	}

	return r
}
