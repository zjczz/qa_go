package serializer

import "likezh/model"

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
		user, _ := model.GetUser(question.UserId)
		questionData := QuestionData{
			ID:        question.ID,
			Nickname:  user.UserProfile.Nickname,
			Title:     question.Title,
			Content:   question.Content,
			Avatar:    user.UserProfile.Avatar,
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
