package v1

import (
	"qa_go/model"
	"qa_go/serializer"
)

// EditQuestionService 管理修改问题的服务
type EditQuestionService struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content"`
}

// isQuestionOwner 判断用户是否拥有该问题
func isQuestionOwner(user *model.User, id uint) *serializer.Response {
	question, err := model.GetQuestion(id)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeQuestionIdError)
	}
	if question.UserID != user.ID {
		return serializer.ErrorResponse(serializer.CodeQuestionNotOwn)
	}
	return nil
}

// EditQuestion 修改问题
func (editQuestionService *EditQuestionService) EditQuestion(user *model.User, id uint) *serializer.Response {
	if err := isQuestionOwner(user, id); err != nil {
		return err
	}
	if question, err := model.UpdateQuestion(id, map[string]interface{}{
		"title":   editQuestionService.Title,
		"content": editQuestionService.Content,
	}); err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	} else {
		return serializer.OkResponse(serializer.BuildQuestionResponse(&question))
	}
}

// DeleteQuestion 删除问题
func DeleteQuestion(user *model.User, id uint) *serializer.Response {
	if err := isQuestionOwner(user, id); err != nil {
		return err
	}
	if err := model.DeleteQuestion(id); err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	} else {
		return serializer.OkResponse(nil)
	}
}
