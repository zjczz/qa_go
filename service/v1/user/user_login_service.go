package v1

import (
	"qa_go/middleware/auth"
	"qa_go/model"
	"qa_go/serializer"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=18"`
}

func GenerateToken(userId uint) (string, error) {
	claim := auth.JwtClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(auth.JwtExpireTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(auth.JwtSecretKey)
}

// Login 用户登录函数
func (service *UserLoginService) Login() *serializer.Response {
	var user model.User

	if err := model.DB.Where("username = ?", service.Username).Preload("UserProfile").First(&user).Error; err != nil {
		return serializer.ErrorResponse(serializer.CodeUserNotExistError)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ErrorResponse(serializer.CodePasswordError)
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeUnknownError)
	}

	data := gin.H{
		"token": token,
		"user":  serializer.BuildUserData(&user),
	}
	return serializer.OkResponse(&data)
}
