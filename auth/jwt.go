package auth

import (
	"github.com/dgrijalva/jwt-go"
	"qa_go/conf"
	"time"
)

const (
	JwtExpireTime = time.Hour * 24
)

var (
	JwtSecretKey = conf.JwtSecretKey
)

// Jwt 编码的结构体
type JwtClaim struct {
	jwt.StandardClaims
	UserId uint
}
