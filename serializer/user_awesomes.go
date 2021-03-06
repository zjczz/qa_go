package serializer

import "qa_go/model"

// 点赞列表的每一项数据
type AwesomesData struct {
	ID          uint   `json:"id"`
	QuestionID  uint   `json:"qid"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	LikeCount   uint   `json:"like_count"`
	LikeStatus  uint   `json:"like_status"`
}

// 点赞列表响应信息
type AwesomesResponse struct {
	Count   int            `json:"count"`
	Answers []AwesomesData `json:"answers"`
}

// 序列化点赞列表响应
func BuildAwesomesResponse(answers []model.Answer, uid uint) *AwesomesResponse {
	var awesomesResponse AwesomesResponse
	awesomesResponse.Answers = make([]AwesomesData, 0, len(answers))

	for _, answer := range answers {
		userProfile, _ := model.GetUserProfile(answer.UserID)
		likes, _ := model.GetAnswerLikedCount(answer.ID)
		status, _ := model.GetUserLikeStatus(uid, answer.ID)
		question, _ := model.GetQuestion(answer.QuestionID)

		awesomesResponse.Answers = append(awesomesResponse.Answers, AwesomesData{
			ID:          answer.ID,
			QuestionID:  answer.QuestionID,
			Title:       question.Title,
			Content:     answer.Content,
			Avatar:      userProfile.Avatar,
			Nickname:    userProfile.Nickname,
			Description: userProfile.Description,
			CreatedAt:   answer.CreatedAt.Unix(),
			LikeCount:   likes,
			LikeStatus:  status,
		})
	}
	awesomesResponse.Count = len(awesomesResponse.Answers)
	return &awesomesResponse
}
