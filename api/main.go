package api

import (
	"encoding/json"
	"fmt"
	"likezh/conf"
	"likezh/model"
	"likezh/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// Index 主页
func Index(c *gin.Context) {
	c.String(http.StatusOK, "================   Welcome to likezh Restful API Index Page!     https://github.com/L-HeliantHuS/likezh   ================")
}

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Msg: "Pong",
	}.Result())
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return serializer.Response{
				Code:  serializer.UserInputError,
				Msg:   fmt.Sprintf("%s%s", field, tag),
				Error: fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Code:  serializer.UserInputError,
			Msg:   "JSON类型不匹配",
			Error: fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Code:  serializer.UserInputError,
		Msg:   "参数错误",
		Error: fmt.Sprint(err),
	}
}
