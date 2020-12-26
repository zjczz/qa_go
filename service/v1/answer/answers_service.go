package v1

import (
	"qa_go/model"
	"qa_go/serializer"
)

// 获取回答列表
func FindAnswers(questionID uint, orderType int, limit int, offset int) *serializer.Response {

	var answers []model.Answer
	var err error

	switch orderType {
	case 0:
		answers, err = model.GetAnswersByScore(questionID, limit, offset)
	case 1:
		answers, err = model.GetAnswersByTime(questionID, limit, offset)
	default:
		answers, err = model.GetAnswersByScore(questionID, limit, offset)
	}

	if err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}

	return serializer.OkResponse(serializer.BuildAnswersResponse(answers))
}
