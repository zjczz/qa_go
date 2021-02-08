package serializer

import "qa_go/model"

//AnswerData 单个回答信息
type AnswerData struct {
	ID          uint   `json:"id"`
	QuestionID  uint   `json:"qid"`
	UserID      uint   `json:"uid"`
	Content     string `json:"content"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	Own         bool   `json:"own"`
	LikeCount   uint   `json:"like_count"`
	LikeStatus  uint   `json:"like_status"`
}

// BuildAnswer 序列化单个回答
func BuildAnswer(answer *model.Answer, uid uint) *AnswerData {
	profile, _ := model.GetUserProfile(answer.UserID)
	likes, _ := model.GetAnswerLikedCount(answer.ID)
	status, _ := model.GetUserLikeStatus(uid, answer.ID)

	return &AnswerData{
		ID:          answer.ID,
		QuestionID:  answer.QuestionID,
		UserID:      answer.UserID,
		Content:     answer.Content,
		Avatar:      profile.Avatar,
		Nickname:    profile.Nickname,
		Description: profile.Description,
		CreatedAt:   answer.CreatedAt.Unix(),
		Own:         uid == answer.UserID,
		LikeCount:   likes,
		LikeStatus:  status,
	}
}

//AnswerResponse 单个回答响应信息
type AnswerResponse struct {
	Answer *AnswerData `json:"answer"`
}

// BuildAnswerResponse 序列化单个问题响应
func BuildAnswerResponse(answer *model.Answer, uid uint) *AnswerResponse {
	return &AnswerResponse{
		Answer: BuildAnswer(answer, uid),
	}
}
