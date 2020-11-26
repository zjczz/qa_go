package api

import (
	"likezh/model"
	"github.com/gin-gonic/gin"
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

