package v1

import (
	"qa_go/model"
	"qa_go/serializer"
)

// 获取热门问题列表
func FindHotQuestions(limit int, offset int) *serializer.Response {
	if questions, err := model.GetHotQuestions(limit, offset); err == nil {
		return serializer.OkResponse(serializer.BuildHotQuestionsResponse(questions))
	}
	return serializer.ErrorResponse(serializer.CodeDatabaseError)
}

// 获取首页问题列表，并加载其回答
func FindQuestions(limit int, offset int) *serializer.Response {
	if questions, err := model.GetQuestions(limit, offset); err == nil {
		return serializer.OkResponse(serializer.BuildQuestionsResponse(questions))
	}
	return serializer.ErrorResponse(serializer.CodeDatabaseError)
}
