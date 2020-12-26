package serializer

import "qa_go/model"

// UserData 单个用户信息
type UserData struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

// BuildUserData 序列化单个用户
func BuildUserData(user *model.User) *UserData {
	profile, _ := model.GetUserProfile(user.ID)
	return &UserData{
		ID:          user.ID,
		Username:    user.Username,
		Nickname:    profile.Nickname,
		Email:       profile.Email,
		Status:      profile.Status,
		Avatar:      profile.Avatar,
		Description: profile.Description,
		CreatedAt:   user.CreatedAt.Unix(),
	}
}

// UserResponse 单个用户响应信息
type UserResponse struct {
	User *UserData `json:"user"`
}

// BuildUserResponse 序列化单个用户响应
func BuildUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		User: BuildUserData(user),
	}
}
