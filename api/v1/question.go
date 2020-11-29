package v1

import (
	"likezh/api"
	// "likezh/cache"
	"likezh/serializer"
	v1 "likezh/service/v1"

	//"net/http"
	"github.com/gin-gonic/gin"
)

// 发表问题
func QuestionAdd(c *gin.Context) {
	var service v1.QesAddService
	user := api.CurrentUser(c)
	if err := c.ShouldBind(&service); err == nil {
		res := service.QuestionAdd(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))

	}
}
