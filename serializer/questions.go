package serializer

import "qa_go/model"

// 热门问题列表信息
type HotQuestionsData struct {
	Count     int            `json:"count"`
	Questions []QuestionData `json:"questions"`
}

// 首页问题列表信息
type QuestionsData struct {
	Count     int             `json:"count"`
	Questions []QuestionBrief `json:"questions"`
}

type QuestionBrief struct {
	ID     uint         `json:"id"`
	Title  string       `json:"title"`
	Answer *AnswerBrief `json:"answer"`
}

// 单个回答信息
type AnswerBrief struct {
	ID          uint   `json:"id"`
	Content     string `json:"content"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
}

// 序列化热门问题列表
func BuildHotQuestions(questions []model.Question) *HotQuestionsData {
	questionsData := HotQuestionsData{}
	questionsData.Count = len(questions)
	questionsData.Questions = make([]QuestionData, len(questions))
	for index, question := range questions {
		profile, _ := model.GetUserProfile(question.UserID)
		questionData := QuestionData{
			ID:        question.ID,
			Nickname:  profile.Nickname,
			Title:     question.Title,
			Content:   question.Content,
			Avatar:    profile.Avatar,
			CreatedAt: question.CreatedAt.Unix(),
		}
		questionsData.Questions[index] = questionData
	}
	return &questionsData
}

// 序列化首页问题列表
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

// 热门问题列表响应信息
type HotQuestionsResponse struct {
	HotQuestionsData
}

// 首页问题列表响应信息
type QuestionsResponse struct {
	QuestionsData
}

// 序列化热门问题列表响应
func BuildHotQuestionsResponse(questions []model.Question) *HotQuestionsResponse {
	return &HotQuestionsResponse{
		*BuildHotQuestions(questions),
	}
}

// 序列化首页问题列表响应
func BuildQuestionsResponse(questions []model.Question) *QuestionsResponse {
	return &QuestionsResponse{
		*BuildQuestions(questions),
	}
}
