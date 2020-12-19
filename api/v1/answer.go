package v1

import (
	"qa_go/serializer"
	"strconv"

	v1 "qa_go/service/v1"

	"qa_go/api"

	"github.com/gin-gonic/gin"
)

func AddAnswer(c *gin.Context) {
	// qid 所属问题id
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	// 解析参数
	var service v1.AddAnswerService
	err = c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}

	user := api.CurrentUser(c)
	res := service.AddAnswer(user, uint(qid))
	c.JSON(200, res)
}

func FindAnswer(c *gin.Context) {
	// qid 所属问题id
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	// aid 回答id
	aid, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}

	res := v1.FindOneAnswer(uint(qid), uint(aid))
	c.JSON(200, res)
}
