package serializer

import "qa_go/model"

// 单个回答信息
type AnswerData struct {
	ID          uint   `json:"id"`
	QuestionID  uint   `json:"qid"`
	Content     string `json:"content"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

// 序列化单个回答
func BuildAnswer(ans *model.Answer) *AnswerData {
	user, _ := model.GetUser(ans.UserID)
	return &AnswerData{
		ID:          ans.ID,
		QuestionID:  ans.QuestionID,
		Content:     ans.Content,
		Avatar:      user.UserProfile.Avatar,
		Nickname:    user.UserProfile.Nickname,
		Description: user.UserProfile.Description,
		CreatedAt:   ans.CreatedAt.Unix(),
	}
}

// 单个回答响应信息
type AnswerResponse struct {
	Answer *AnswerData `json:"answer"`
}

// 序列化单个问题响应
func BuildAnswerResponse(answer *model.Answer) *AnswerResponse {
	return &AnswerResponse{
		Answer: BuildAnswer(answer),
	}
}
