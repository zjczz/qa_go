package api

import (
	"likezh/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 主页
func Index(c *gin.Context) {
	c.String(http.StatusOK, "======= https://github.com/zjczz/likezh =======")
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if userID, _ := c.Get("user"); userID != nil {
		if user,err:=model.GetUser(userID);err==nil{
			return &user
		}
	}
	return nil
}
