package v1

import (
	"qa_go/model"
	"qa_go/serializer"
)

func FindOneQuestion(id uint,uid uint) *serializer.Response {
	if question, err := model.GetQuestion(id); err == nil {
		return serializer.OkResponse(serializer.BuildQesResponse(&question,uid))
	}
	return serializer.ErrorResponse(serializer.CodeQuestionIdError)
}
