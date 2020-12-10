package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qa_go/api"
	"qa_go/cache"
	"qa_go/serializer"
	v1 "qa_go/service/v1"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var service v1.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var service v1.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := api.CurrentUser(c)
	c.JSON(http.StatusOK, serializer.OkResponse(serializer.BuildUserResponse(user)))
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	user := api.CurrentUser(c)
	var service v1.ChangePassword
	if err := c.ShouldBind(&service); err == nil {
		res := service.Change(user)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.ErrorResponse(serializer.CodeParamError))
	}
}

// Logout 用户退出登录
func Logout(c *gin.Context) {
	token, _ := c.Get("token")
	tokenString := token.(string)

	cache.RedisClient.SAdd("jwt:baned", tokenString)
	c.JSON(http.StatusOK, serializer.Response{
		Msg: "已退出登录",
	})
}
