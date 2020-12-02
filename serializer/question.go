package serializer

import "likezh/model"

// 单个问题信息
type QuestionData struct {
	ID        uint    `json:"id"`
	Title  	  string  `json:"title"`
	Content   string  `json:"content"`
	Nickname  string  `json:"nickname"`
	Avatar    string  `json:"avatar"`
	CreatedAt int64   `json:"created_at"`
}

// 序列化单个问题
func BuildQuestion(qes *model.Question) *QuestionData {
	user,_:=model.GetUser(qes.UserId)
	return &QuestionData{
		ID:        qes.ID,
		Nickname:  user.UserProfile.Nickname,
		Title: qes.Title,
		Content: qes.Content,
		Avatar:    user.UserProfile.Avatar,
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
