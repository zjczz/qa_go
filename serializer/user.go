package serializer

import "likezh/model"

// userdata 用户序列化器
type userdata struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Status    int    `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// buildUser 序列化用户
func buildUser(user model.User) *userdata {
	return &userdata{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// UserResponse 单个用户序列化
type UserResponse struct {
	User *userdata `json:"user"`
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) *UserResponse {
	return &UserResponse{
		User: buildUser(user),
	}
}
