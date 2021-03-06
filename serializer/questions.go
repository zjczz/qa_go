package serializer

import (
	"qa_go/cache"
	"qa_go/model"
	"strconv"
)

// Redis 缓存的一条热门问题
type CacheHotQuestion struct {
	ID    string
	Title string
}

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
func BuildQuestions(questions []CacheHotQuestion) *QuestionsData {
	questionsData := QuestionsData{}
	questionsData.Count = len(questions)
	questionsData.Questions = make([]QuestionBrief, len(questions))
	for index, question := range questions {
		qid, _ := strconv.Atoi(question.ID)
		questionData := QuestionBrief{
			ID:     uint(qid),
			Title:  question.Title,
			Answer: nil,
		}
		if cache.RedisClient.HExists(cache.KeyHotAnswer, question.ID).Val() {
			aid, _ := cache.RedisClient.HGet(cache.KeyHotAnswer, question.ID).Int()
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
		//if answer := model.GetHotAnswer(question.ID); answer != nil {
		//	profile, _ := model.GetUserProfile(answer.UserID)
		//	likes, _ := model.GetAnswerLikedCount(answer.ID)
		//	answerBrief := AnswerBrief{}
		//	answerBrief.ID = answer.ID
		//	answerBrief.Content = answer.Content
		//	answerBrief.Avatar = profile.Avatar
		//	answerBrief.Nickname = profile.Nickname
		//	answerBrief.Description = profile.Description
		//	answerBrief.LikeCount = likes
		//	questionData.Answer = &answerBrief
		//}
		questionsData.Questions[index] = questionData
	}
	return &questionsData
}

// 首页推荐列表响应信息
type QuestionsResponse struct {
	QuestionsData
}

// 序列化首页推荐列表响应
func BuildQuestionsResponse(questions []CacheHotQuestion) *QuestionsResponse {
	return &QuestionsResponse{
		*BuildQuestions(questions),
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
func BuildHotQuestionsResponse(questions []HotQuestionsData) *HotQuestionsResponse {
	return &HotQuestionsResponse{
		Count:     len(questions),
		Questions: questions,
	}
}
