package serializer

import "qa_go/model"

// 单个问题
type QuestionData struct {
	ID          uint   `json:"id"`
	UID         uint   `json:"uid"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedAt   int64  `json:"created_at"`
	Own         bool   `json:"own"`
	AnswerCount int64  `json:"answer_count"`
}

// 序列化单个问题
func BuildQuestion(question *model.Question, uid uint) *QuestionData {
	count := model.GetAnswerCount(question.ID)
	return &QuestionData{
		ID:          question.ID,
		UID:         question.UserID,
		Title:       question.Title,
		Content:     question.Content,
		CreatedAt:   question.CreatedAt.Unix(),
		Own:         uid == question.UserID,
		AnswerCount: count,
	}
}

// 单个问题响应信息
type QuestionResponse struct {
	Question *QuestionData `json:"question"`
}

// 序列化单个普通问题响应
func BuildQuestionResponse(question *model.Question, uid uint) *QuestionResponse {
	return &QuestionResponse{
		Question: BuildQuestion(question, uid),
	}
}
