package serializer

import "qa_go/model"

// 单个热点问题信息
type QuestionData struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}
//单个普通问题
type QesData struct{
	ID        uint   `json:"id"`
	UID  	  uint   `json:"uid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	Own    	  bool	 `json:"own"`
	AnswerCount int64  `json:"answer_count"`
}
// 序列化热点单个问题
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
// 序列化普通单个问题
func BuildQes(qes *model.Question,uid uint) *QesData {
	cnt:=model.GetAnswerByQesID(qes.ID)
	// own:=false
	// if uid==qes.UserID{
	// 	own=true
	// }
	return &QesData{
		ID:        qes.ID,
		UID: 	   qes.UserID,
		Title:     qes.Title,
		Content:   qes.Content,
		CreatedAt: qes.CreatedAt.Unix(),
		Own:       uid==qes.UserID,
		AnswerCount:cnt ,
	}
}
// 单个热点问题响应信息
type QuestionResponse struct {
	Question *QuestionData `json:"question"`
}
//单个普通问题响应信息
type QesResponse struct{
	Question *QesData `json:"question"`
}
// 序列化单个热点问题响应
func BuildQuestionResponse(question *model.Question) *QuestionResponse {
	return &QuestionResponse{
		Question: BuildQuestion(question),
	}
}
//序列化单个普通问题响应
func BuildQesResponse(qes *model.Question,uid uint) *QesResponse{
	return &QesResponse{
		Question:BuildQes(qes,uid),
	}
}