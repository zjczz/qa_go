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
	{
		// 注册
		v1Group.POST("/user/register", v1.UserRegister)
		// 登录
		v1Group.POST("/user/login", v1.UserLogin)

		// 获取首页推荐列表
		v1Group.GET("/questions", v1.FindQuestions)
		// 获取问题热榜
		v1Group.GET("/hot_questions", v1.FindHotQuestions)
		// 获取回答列表
		v1Group.GET("/questions/:qid/answers", v1.FindAnswers)

		// 可选token
		jwtSelect := v1Group.Group("")
		jwtSelect.Use(auth.JwtWithAnonymous())
		{
			// 查看单个问题
			jwtSelect.GET("/questions/:qid", v1.FindOneQuestion)
			// 查看单个回答
			jwtSelect.GET("/questions/:qid/answers/:aid", v1.FindAnswer)
		}

		// 需要登录权限
		jwt := v1Group.Group("")
		jwt.Use(auth.JwtRequired())
		{
			// 查看个人信息
			jwt.GET("/user/me", v1.UserMe)
			// 退出登录
			jwt.POST("/user/logout", v1.Logout)
			// 查看个人发布问题
			jwt.GET("/user/questions", v1.GetUserQuestions)
			// 查看个人发布回答
			jwt.GET("/user/answers", v1.GetUserAnswers)
			// 查看点赞回答列表
			jwt.GET("/user/awesomes", v1.Awesomes)

			// 发布问题
			jwt.POST("/questions", v1.QuestionAdd)
			// 修改问题
			jwt.PUT("/questions/:qid", v1.EditQuestion)
			// 删除问题
			jwt.DELETE("/questions/:qid", v1.DeleteQuestion)

			// 回答问题
			jwt.POST("/questions/:qid/answers", v1.AddAnswer)
			// 修改回答
			jwt.PUT("/questions/:qid/answers/:aid", v1.UpdateAnswer)
			// 删除回答
			jwt.DELETE("/questions/:qid/answers/:aid", v1.DeleteAnswer)
			// 点赞回答
			jwt.POST("/answers/:aid/voters", v1.Voter)
		}
	}

	return r
}
