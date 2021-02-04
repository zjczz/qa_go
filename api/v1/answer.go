package v1

import (
	"qa_go/api"
	"qa_go/serializer"
	v1 "qa_go/service/v1/answer"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加回答
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

// 查看回答
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
	user := api.CurrentUser(c)
	var uid uint
	if user == nil {
		uid = 0
	} else {
		uid = user.ID
	}
	res := v1.FindOneAnswer(uint(qid), uint(aid), uid)
	c.JSON(200, res)
}

// 修改回答
func UpdateAnswer(c *gin.Context) {
	qid, err1 := strconv.Atoi(c.Param("qid"))
	aid, err2 := strconv.Atoi(c.Param("aid"))
	if err1 != nil || err2 != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	var service v1.UpdateAnswerService
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	user := api.CurrentUser(c)
	res := service.UpdateAnswer(user, uint(qid), uint(aid))
	c.JSON(200, res)
}

// 删除回答
func DeleteAnswer(c *gin.Context) {
	qid, err1 := strconv.Atoi(c.Param("qid"))
	aid, err2 := strconv.Atoi(c.Param("aid"))
	if err1 != nil || err2 != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}

	user := api.CurrentUser(c)
	res := v1.DeleteAnswer(user, uint(qid), uint(aid))
	c.JSON(200, res)
}

func FindAnswers(c *gin.Context) {
	qid, err1 := strconv.Atoi(c.Param("qid"))
	orderType, err2 := strconv.Atoi(c.DefaultQuery("type", "0"))
	limit, err3 := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, err4 := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
	}
	res := v1.FindAnswers(uint(qid), orderType, limit, offset)
	c.JSON(200, res)
}

//点赞
func Voter(c *gin.Context) {
	aid, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
	}
	// 解析参数
	var status v1.VoterService
	err = c.ShouldBind(&status)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	user := api.CurrentUser(c)
	res := v1.Voter(user.ID, uint(aid), status.Type)
	c.JSON(200, res)
}

//获取用户赞过的回答
func Awesomes(c *gin.Context) {
	user := api.CurrentUser(c)
	res := v1.GetAwesomes(user.ID)
	c.JSON(200, res)
}
