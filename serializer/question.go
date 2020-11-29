package serializer
import "likezh/model"


// 序列化器
type questiondata struct {
	ID        uint   `json:"id"`
	Nickname  string `json:"nickname"`
	Title  	  string  `json:"title"`
	Content   string  `json:"content"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

//序列化
func buildQuestion(qes model.Question) *questiondata {
	user,_:=model.GetUser(qes.UserId)
	return &questiondata{
		ID:        qes.ID,
		Nickname:  user.Nickname,
		Title: qes.Title,
		Content: qes.Content,
		Avatar:    user.Avatar,
		CreatedAt: qes.CreatedAt.Unix(),
	}
}

//
type QuestionResponse struct {
	Question *questiondata `json:"question"`
}

//
func BuildQuestionResponse(question model.Question) *QuestionResponse {
	return &QuestionResponse{
		Question: buildQuestion(question),
	}
}
