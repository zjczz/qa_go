package serializer

import "qa_go/model"

// 单个问题信息
type QuestionData struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// 序列化单个问题
func BuildQuestion(qes *model.Question) *QuestionData {
	profile, _ := model.GetUserProfile(qes.UserID)
	return &QuestionData{
		ID:        qes.ID,
		Nickname:  profile.Nickname,
		Title:     qes.Title,
		Content:   qes.Content,
		Avatar:    profile.Avatar,
		CreatedAt: qes.CreatedAt.Unix(),
	}
}

// 单个问题响应信息
type QuestionResponse struct {
	Question *QuestionData `json:"question"`
}

// 序列化单个问题响应
func BuildQuestionResponse(question *model.Question) *QuestionResponse {
	return &QuestionResponse{
		Question: BuildQuestion(question),
	}
}
