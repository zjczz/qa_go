package serializer

import "likezh/model"

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
    return &UserData{
        ID:        user.ID,
        Username:  user.Username,
        Nickname:  user.UserProfile.Nickname,
        Email: user.UserProfile.Email,
        Status:    user.UserProfile.Status,
        Avatar:    user.UserProfile.Avatar,
        Description: user.UserProfile.Description,
        CreatedAt: user.CreatedAt.Unix(),
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
