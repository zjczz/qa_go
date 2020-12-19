package v1

import (
	"qa_go/model"
	"qa_go/serializer"
)

// AddAnswerService 管理回答问题的服务
type AddAnswerService struct {
	Content string `form:"content" json:"content" binding:"required"`
}

// 回答问题
func (service *AddAnswerService) AddAnswer(user *model.User, qid uint) *serializer.Response {
	answer := &model.Answer{
		UserID:     user.ID,
		QuestionID: qid,
		Content:    service.Content,
	}

	if err := model.DB.Create(answer).Error; err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	return serializer.OkResponse(serializer.BuildAnswerResponse(answer))
}

// 查看单个问题
func FindOneAnswer(qid uint, aid uint) *serializer.Response {
	if answer, err := model.GetAnswer(aid); err == nil {
		if answer.QuestionID != qid {
			return serializer.ErrorResponse(serializer.CodeQidMismatchError)
		}
		return serializer.OkResponse(serializer.BuildAnswerResponse(answer))
	}
	return serializer.ErrorResponse(serializer.CodeAnswerIdError)
}
