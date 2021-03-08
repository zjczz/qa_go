package serializer

import (
	"qa_go/cache"
	"qa_go/model"
	"strconv"
)

// 首页推荐列表单个问题
type QuestionBrief struct {
	ID     uint         `json:"id"`
	Title  string       `json:"title"`
	Answer *AnswerBrief `json:"answer"`
}

// 首页推荐列表单个回答
type AnswerBrief struct {
	ID          uint   `json:"id"`
	Content     string `json:"content"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
	LikeCount   uint   `json:"like_count"`
}

// 首页推荐列表
type QuestionsData struct {
	Count     int             `json:"count"`
	Questions []QuestionBrief `json:"questions"`
}

// 序列化首页推荐列表
func BuildQuestions(questionsID []string) *QuestionsData {
	questionsData := QuestionsData{}
	questionsData.Questions = make([]QuestionBrief, 0, len(questionsID))
	for _, questionID := range questionsID {
		qid, _ := strconv.Atoi(questionID)
		if title, err := cache.RedisClient.HGet(cache.KeyHotQuestionTitle, questionID).Result(); err == nil {
			questionData := QuestionBrief{
				ID:     uint(qid),
				Title:  title,
				Answer: nil,
			}
			if aidStr, err := cache.RedisClient.HGet(cache.KeyHotAnswer, questionID).Result(); err == nil {
				aid, _ := strconv.Atoi(aidStr)
				if answer, err := model.GetAnswer(uint(aid)); err == nil {
					profile, _ := model.GetUserProfile(answer.UserID)
					likes, _ := model.GetAnswerLikedCount(answer.ID)
					answerBrief := AnswerBrief{}
					answerBrief.ID = answer.ID
					answerBrief.Content = answer.Content
					answerBrief.Avatar = profile.Avatar
					answerBrief.Nickname = profile.Nickname
					answerBrief.Description = profile.Description
					answerBrief.LikeCount = likes
					questionData.Answer = &answerBrief
				}
			}
			questionsData.Questions = append(questionsData.Questions, questionData)
		}
	}
	questionsData.Count = len(questionsData.Questions)
	return &questionsData
}

// 首页推荐列表响应信息
type QuestionsResponse struct {
	QuestionsData
}

// 序列化首页推荐列表响应
func BuildQuestionsResponse(questionsID []string) *QuestionsResponse {
	return &QuestionsResponse{
		*BuildQuestions(questionsID),
	}
}

// 单个热点问题信息
type HotQuestionsData struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Hot   uint   `json:"hot"`
}

// 热门问题列表信息
type HotQuestionsResponse struct {
	Count     int                `json:"count"`
	Questions []HotQuestionsData `json:"questions"`
}

// 序列化热门问题列表响应
func BuildHotQuestionsResponse(questions []string) *HotQuestionsResponse {
	response := HotQuestionsResponse{}
	response.Questions = make([]HotQuestionsData, 0, len(questions))
	for _, questionsID := range questions {
		qid, _ := strconv.Atoi(questionsID)
		score, err := cache.RedisClient.ZScore(cache.KeyHotQuestions, questionsID).Result()
		if err != nil {
			continue
		}
		if title, err := cache.RedisClient.HGet(cache.KeyHotQuestionTitle, questionsID).Result(); err == nil {
			response.Questions = append(response.Questions, HotQuestionsData{
				ID:    uint(qid),
				Title: title,
				Hot:   uint(score),
			})
		}
	}
	response.Count = len(response.Questions)
	return &response
}
