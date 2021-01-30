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
	Likes       uint  `json:"likes"`
	UserLikeStatus uint `json:"userlikestatus"`
	}

// BuildAnswer 序列化单个回答
func BuildAnswer(answer *model.Answer, uid uint) *AnswerData {
	profile, _ := model.GetUserProfile(answer.UserID)
	likes,_:=model.GetAnswerlikedCount(answer.ID)
	status,_:=model.GetUserLikeStatus(uid,answer.ID)
	
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
		Likes:		 likes ,
		UserLikeStatus: status,
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

//AwsResponse 回答列表响应信息
type AwsResponse struct {
	Count   int           `json:"count"`
	Answers []AnswerData `json:"answers"`
}

//BuildAwsResponse 序列化回答列表响应
func BuildAwsResponse(answers []model.Answer,uid uint) *AwsResponse {
	var answersResponse AwsResponse
	answersResponse.Count = len(answers)
	
	for _, answer := range answers {
		userProfile, _ := model.GetUserProfile(answer.UserID)
		likes,_:=model.GetAnswerlikedCount(answer.ID)
		status,_:=model.GetUserLikeStatus(uid,answer.ID)
		answersResponse.Answers = append(answersResponse.Answers, AnswerData{
			ID:          answer.ID,
			QuestionID:  answer.QuestionID,
			UserID:      answer.UserID,
			Content:     answer.Content,
			Avatar:      userProfile.Avatar,
			Nickname:    userProfile.Nickname,
			Description: userProfile.Description,
			CreatedAt:   answer.CreatedAt.Unix(),
			Own:         uid == answer.UserID,
			Likes:		 likes ,
			UserLikeStatus: status,
		})
	}
	return &answersResponse
}