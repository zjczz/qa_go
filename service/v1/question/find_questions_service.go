package v1

import (
	"qa_go/cache"
	"qa_go/model"
	"qa_go/serializer"
	"strconv"
	"strings"
)

// 获取首页问题列表，并加载其回答
func FindQuestions(limit int, offset int) *serializer.Response {
	if questions, err := model.GetQuestions(limit, offset); err == nil {
		return serializer.OkResponse(serializer.BuildQuestionsResponse(questions))
	}
	return serializer.ErrorResponse(serializer.CodeDatabaseError)
}

// 获取热门问题列表
func FindHotQuestions() *serializer.Response {
	var questionsCache []string
	if err := cache.RedisClient.ZRevRange(cache.KeyHotQuestions, 0, 49).ScanSlice(&questionsCache); err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}

	var hotQuestions []serializer.HotQuestionsData
	for _, member := range questionsCache {
		score, _ := cache.RedisClient.ZScore(cache.KeyHotQuestions, member).Result()
		splitK := strings.Split(member, ":")
		id, _ := strconv.Atoi(splitK[0])
		title := splitK[1]
		hotQuestions = append(hotQuestions, serializer.HotQuestionsData{
			ID:    uint(id),
			Title: title,
			Hot:   uint(score),
		})
	}
	return serializer.OkResponse(serializer.BuildHotQuestionsResponse(hotQuestions))
}
