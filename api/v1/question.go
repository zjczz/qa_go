package v1

import (
	"likezh/api"
	// "likezh/cache"
	"likezh/serializer"
	v1 "likezh/service/v1"

	//"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
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

// 查看单个问题
func FindOneQuestion(c *gin.Context) {
	qid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	res := v1.FindOneQuestion(uint(qid))
	c.JSON(200, res)
}

// 获取问题列表
func FindQuestions(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	res := v1.FindQuestions(limit, offset)
	c.JSON(200, res)
}
