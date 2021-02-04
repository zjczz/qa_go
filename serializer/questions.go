package serializer

import (
	"qa_go/model"
	"sort"
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
func BuildQuestions(questions []model.Question) *QuestionsData {
	questionsData := QuestionsData{}
	questionsData.Count = len(questions)
	questionsData.Questions = make([]QuestionBrief, len(questions))
	for index, question := range questions {

		questionData := QuestionBrief{
			ID:     question.ID,
			Title:  question.Title,
			Answer: nil,
		}
		if len(question.Answers) != 0 {
			profile, _ := model.GetUserProfile(question.Answers[0].UserID)
			likes, _ := model.GetAnswerlikedCount(question.Answers[0].ID)
			answer := AnswerBrief{}
			answer.ID = question.Answers[0].ID
			answer.Content = question.Answers[0].Content
			answer.Avatar = profile.Avatar
			answer.Nickname = profile.Nickname
			answer.Description = profile.Description
			answer.LikeCount = likes
			questionData.Answer = &answer
		}
		questionsData.Questions[index] = questionData
	}
	return &questionsData
}

// 首页推荐列表响应信息
type QuestionsResponse struct {
	QuestionsData
}

// 序列化首页推荐列表响应
func BuildQuestionsResponse(questions []model.Question) *QuestionsResponse {
	return &QuestionsResponse{
		*BuildQuestions(questions),
	}
}

// 单个热点问题信息
type HotQuestionData struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Hot   uint   `json:"hot"`
}

// 热门问题列表信息
type HotQuestionsData struct {
	Count     int               `json:"count"`
	Questions []HotQuestionData `json:"questions"`
}

// 自定义类型，便于实现自定义排序方法
type HotQuestions []HotQuestionData

// 实现排序接口方法
func (q HotQuestions) Len() int {
	return len(q)
}
func (q HotQuestions) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q HotQuestions) Less(i, j int) bool {
	return q[i].Hot > q[j].Hot
}

// 序列化热门问题列表
func BuildHotQuestions(questions []model.Question) *HotQuestionsData {
	hotQuestionsData := HotQuestionsData{}
	hotQuestionsData.Count = len(questions)
	var hots HotQuestions = make([]HotQuestionData, len(questions))
	for index, question := range questions {
		hotQuestionData := HotQuestionData{
			ID:    question.ID,
			Title: question.Title,
			Hot:   question.AnswerCount,
		}
		hots[index] = hotQuestionData
	}
	// 按热度值排序
	sort.Sort(hots)
	hotQuestionsData.Questions = hots
	return &hotQuestionsData
}

// 热门问题列表响应信息
type HotQuestionsResponse struct {
	HotQuestionsData
}

// 序列化热门问题列表响应
func BuildHotQuestionsResponse(questions []model.Question) *HotQuestionsResponse {
	return &HotQuestionsResponse{
		*BuildHotQuestions(questions),
	}
}
