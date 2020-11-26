package v1

import (
	"likezh/auth"
	"likezh/conf"
	"likezh/model"
	"likezh/serializer"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=18"`
}

func GenerateToken(user model.User, ExpiresTime int64) (string, error) {
	claims := auth.Jwt{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresTime,
			IssuedAt:  time.Now().Unix(),
		},
		Data: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, err := token.SignedString(conf.SigningKey)
	return jwtString, err
}

// Login 用户登录函数
func (service *UserLoginService) Login() *serializer.Response {
	var user model.User
	ExpiresTime := time.Now().Add(time.Hour * time.Duration(720)).Unix()

	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err != nil {
		return serializer.ErrorResponse(serializer.CodeUserNotExistError)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ErrorResponse(serializer.CodePasswordError)
	}

	token, err := GenerateToken(user, ExpiresTime)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeUnknownError)
	}

	data := gin.H{
		"token": token,
	}
	return serializer.OkResponse(data)
}
