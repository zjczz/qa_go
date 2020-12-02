package v1

import (
	"likezh/model"
	"likezh/serializer"
)

func FindQuestions(limit int, offset int) *serializer.Response {
	if questions, err := model.GetQuestions(limit, offset); err == nil {
		return serializer.OkResponse(serializer.BuildQuestionsResponse(questions))
	}
	return serializer.ErrorResponse(serializer.CodeQuestionsError)
}
