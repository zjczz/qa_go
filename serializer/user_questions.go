package serializer

import "qa_go/model"

// 个人问题列表的每一项数据
type UserQuestionsData struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	AnswerCount uint   `json:"answer_count"`
	CreatedAt   int64  `json:"created_at"`
}

// 个人问题列表响应信息
type UserQuestionsResponse struct {
	Count     int                 `json:"count"`
	Questions []UserQuestionsData `json:"questions"`
}

// 序列化个人问题列表响应
func BuildUserQuestionsResponse(questions []model.Question) *UserQuestionsResponse {
	var userQuestionsResponse UserQuestionsResponse
	userQuestionsResponse.Questions = make([]UserQuestionsData, 0, len(questions))

	for _, question := range questions {
		userQuestionsResponse.Questions = append(userQuestionsResponse.Questions, UserQuestionsData{
			ID:          question.ID,
			Title:       question.Title,
			Content:     question.Content,
			AnswerCount: question.AnswerCount,
			CreatedAt:   question.CreatedAt.Unix(),
		})
	}
	userQuestionsResponse.Count = len(userQuestionsResponse.Questions)
	return &userQuestionsResponse
}
