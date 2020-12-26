package serializer

import "qa_go/model"

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
			answer := AnswerBrief{}
			answer.ID = question.Answers[0].ID
			answer.Content = question.Answers[0].Content
			answer.Avatar = profile.Avatar
			answer.Nickname = profile.Nickname
			answer.Description = profile.Description
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
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// 热门问题列表信息
type HotQuestionsData struct {
	Count     int               `json:"count"`
	Questions []HotQuestionData `json:"questions"`
}

// 序列化热门问题列表
func BuildHotQuestions(questions []model.Question) *HotQuestionsData {
	hotQuestionsData := HotQuestionsData{}
	hotQuestionsData.Count = len(questions)
	hotQuestionsData.Questions = make([]HotQuestionData, len(questions))
	for index, question := range questions {
		profile, _ := model.GetUserProfile(question.UserID)
		hotQuestionData := HotQuestionData{
			ID:        question.ID,
			Nickname:  profile.Nickname,
			Title:     question.Title,
			Content:   question.Content,
			Avatar:    profile.Avatar,
			CreatedAt: question.CreatedAt.Unix(),
		}
		hotQuestionsData.Questions[index] = hotQuestionData
	}
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
