package api

import (
	"encoding/json"
	"fmt"
	"likezh/conf"
	"likezh/model"
	"likezh/serializer"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// Index 主页
func Index(c *gin.Context) {
	c.String(http.StatusOK, "======= https://github.com/zjczz/likezh =======")
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
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Code:  serializer.UserInputError,
			Msg:   "JSON类型不匹配",
		}
	}

	return serializer.Response{
		Code:  serializer.UserInputError,
		Msg:   "参数错误",
	}
}
