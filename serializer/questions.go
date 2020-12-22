package serializer

import "qa_go/model"

// 问题列表信息
type QuestionsData struct {
	Count     int            `json:"count"`
	Questions []QuestionData `json:"questions"`
}

// 序列化问题列表
func BuildQuestions(questions []model.Question) *QuestionsData {
	questionsData := QuestionsData{}
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

// 问题列表响应信息
type QuestionsResponse struct {
	QuestionsData
}

// 序列化问题列表响应
func BuildQuestionsResponse(questions []model.Question) *QuestionsResponse {
	return &QuestionsResponse{
		*BuildQuestions(questions),
	}
}
