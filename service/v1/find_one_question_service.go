package v1

import (
	"qa_go/model"
	"qa_go/serializer"
)

func FindOneQuestion(id uint) *serializer.Response {
	if question, err := model.GetQuestion(id); err == nil {
		return serializer.OkResponse(serializer.BuildQuestionResponse(&question))
	}
	return serializer.ErrorResponse(serializer.CodeQuestionIdError)
}
