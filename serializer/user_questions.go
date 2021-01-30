package serializer

import "qa_go/model"

// 个人问题列表的每一项数据
type UserQuestionsData struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

// 个人问题列表响应信息
type UserQuestionsResponse struct {
	Count     int                 `json:"count"`
	Questions []UserQuestionsData `json:"questions"`
}

// 序列化个人问题列表响应
func BuildUserQuestionsResponse(questions []model.Question) *UserQuestionsResponse {
	var userQuestionsResponse UserQuestionsResponse
	userQuestionsResponse.Count = len(questions)
	for _, question := range questions {
		userQuestionsResponse.Questions = append(userQuestionsResponse.Questions, UserQuestionsData{
			ID:        question.ID,
			Title:     question.Title,
			Content:   question.Content,
			CreatedAt: question.CreatedAt.Unix(),
		})
	}
	return &userQuestionsResponse
}
