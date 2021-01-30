package v1

import (
	"qa_go/model"
	"qa_go/serializer"
)

// AddAnswerService 管理回答问题的服务
type AddAnswerService struct {
	Content string `form:"content" json:"content" binding:"required"`
}

// UpdateAnswerService 管理修改回答
type UpdateAnswerService struct {
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
	return serializer.OkResponse(serializer.BuildAnswerResponse(answer, user.ID))
}

// VoterService 点赞状态
type VoterService struct {
	Type string `form:"type" json:"type" binding:"required"`
}

// 查看单个回答
func FindOneAnswer(qid uint, aid uint, uid uint) *serializer.Response {
	if answer, err := model.GetAnswer(aid); err == nil {
		if answer.QuestionID != qid {
			return serializer.ErrorResponse(serializer.CodeQidMismatchError)
		}
		return serializer.OkResponse(serializer.BuildAnswerResponse(answer, uid))
	}
	return serializer.ErrorResponse(serializer.CodeAnswerIdError)
}

// 修改回答
func (service *UpdateAnswerService) UpdateAnswer(user *model.User, qid uint, aid uint) *serializer.Response {
	_, err := model.GetQuestion(qid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeQuestionIdError)
	}
	answer, err := model.GetAnswer(aid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeAnswerIdError)
	}

	if answer.QuestionID != qid {
		return serializer.ErrorResponse(serializer.CodeQidMismatchError)
	}
	if answer.UserID != user.ID {
		return serializer.ErrorResponse(serializer.CodeAnswerNotOwn)
	}

	answer.Content = service.Content
	if model.DB.Save(answer).Error != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	return serializer.OkResponse(serializer.BuildAnswerResponse(answer, user.ID))
}

// 删除回答
func DeleteAnswer(user *model.User, qid uint, aid uint) *serializer.Response {
	_, err := model.GetQuestion(qid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeQuestionIdError)
	}
	answer, err := model.GetAnswer(aid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeAnswerIdError)
	}

	if answer.QuestionID != qid {
		return serializer.ErrorResponse(serializer.CodeQidMismatchError)
	}
	if answer.UserID != user.ID {
		return serializer.ErrorResponse(serializer.CodeAnswerNotOwn)
	}

	if err := model.DeleteAnswer(aid); err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	} else {
		return serializer.OkResponse(nil)
	}
}

//Voter 点赞
func Voter(uid uint, aid uint, status string) *serializer.Response {
	code := 0
	if status == "up" {
		code = 1
	} else if status == "down" {
		code = 2
	}
	err := model.AddUserLike(uid, aid, uint(code))
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeVoterTypeError)
	}
	return serializer.OkResponse(nil)
}

//获取赞的内容
func GetAwesomes(uid uint) *serializer.Response {
	ids, err := model.GetUserLikes(uid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeAnswerIdError)
	}
	ans, _ := model.GetAnswers(ids)
	return serializer.OkResponse(serializer.BuildAwsResponse(ans, uid))
}
