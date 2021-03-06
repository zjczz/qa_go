package v1

import (
	"qa_go/cache"
	"qa_go/serializer"
)

// 获取首页问题列表，并加载其高赞回答
func FindQuestions(limit int, offset int) *serializer.Response {
	var questionsCache []string
	if err := cache.RedisClient.ZRevRange(cache.KeyHotQuestions, int64(offset), int64(offset+limit-1)).ScanSlice(&questionsCache); err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	return serializer.OkResponse(serializer.BuildQuestionsResponse(questionsCache))
}

// 获取热门问题列表
func FindHotQuestions() *serializer.Response {
	var questionsCache []string
	if err := cache.RedisClient.ZRevRange(cache.KeyHotQuestions, 0, 49).ScanSlice(&questionsCache); err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	return serializer.OkResponse(serializer.BuildHotQuestionsResponse(questionsCache))
}
