package auth

import (
	"net/http"
	"qa_go/cache"
	"qa_go/conf"
	"qa_go/serializer"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	// jwt过期时间
	JwtExpireTime = time.Hour * 720
)

var (
	// jwt加密秘钥
	JwtSecretKey = conf.JwtSecretKey
)

// jwt编码的结构体
type JwtClaim struct {
	jwt.StandardClaims
	UserId uint
}

// jwt中间件，可选匿名不用必须带token
func JwtWithAnonymous() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获得token
		userToken := c.Request.Header.Get("token")
		if userToken == "" {
			return
		}

		// 解码token值
		token, err := jwt.ParseWithClaims(userToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
			return JwtSecretKey, nil
		})
		if err != nil || token.Valid != true {
			return
		}

		// 判断令牌是否已注销
		if result, _ := cache.RedisClient.SIsMember("jwt:baned", token.Raw).Result(); result {
			return
		}

		// 将Token放入Context, 用于退出登录添加黑名单
		c.Set("token", token.Raw)

		// 将token携带的用户id信息存入上下文
		if jwtStruct, ok := token.Claims.(*JwtClaim); ok {
			c.Set("user_id", &jwtStruct.UserId)
		}
	}
}

// jwt中间件，需要在Header中传递token
func JwtRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获得token
		userToken := c.Request.Header.Get("token")
		if userToken == "" {
			// 请求是否携带token
			c.JSON(http.StatusOK, serializer.ErrorResponse(serializer.CodeTokenNotExitError))
			c.Abort()
			return
		}

		// 解码token值
		token, err := jwt.ParseWithClaims(userToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
			return JwtSecretKey, nil
		})
		if err != nil || token.Valid != true {
			// 过期或者非正确
			c.JSON(http.StatusOK, serializer.ErrorResponse(serializer.CodeTokenExpiredError))
			c.Abort()
			return
		}

		// 判断令牌是否已注销
		if result, _ := cache.RedisClient.SIsMember("jwt:baned", token.Raw).Result(); result {
			c.JSON(http.StatusOK, serializer.ErrorResponse(serializer.CodeTokenExpiredError))
			c.Abort()
			return
		}

		// 将Token放入Context, 用于退出登录添加黑名单
		c.Set("token", token.Raw)

		// 将token携带的用户id信息存入上下文
		if jwtStruct, ok := token.Claims.(*JwtClaim); ok {
			c.Set("user_id", &jwtStruct.UserId)
		}
	}
}
