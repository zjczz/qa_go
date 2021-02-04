package serializer

import "qa_go/model"

// 个人回答列表的每一项数据
type UserAnswersData struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"qid"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	LikeCount  uint   `json:"like_count"`
	CreatedAt  int64  `json:"created_at"`
}

// 个人回答列表响应信息
type UserAnswersResponse struct {
	Count   int               `json:"count"`
	Answers []UserAnswersData `json:"answers"`
}

// 序列化个人回答列表响应
func BuildUserAnswersResponse(answers []model.Answer) *UserAnswersResponse {
	var userAnswersResponse UserAnswersResponse
	userAnswersResponse.Count = len(answers)
	for _, answer := range answers {
		question, err := model.GetQuestion(answer.QuestionID)
		if err != nil {
			continue
		}
		likes, _ := model.GetAnswerlikedCount(answer.ID)
		userAnswersResponse.Answers = append(userAnswersResponse.Answers, UserAnswersData{
			ID:         answer.ID,
			QuestionID: answer.QuestionID,
			Title:      question.Title,
			Content:    answer.Content,
			LikeCount:  likes,
			CreatedAt:  answer.CreatedAt.Unix(),
		})
	}
	return &userAnswersResponse
}
