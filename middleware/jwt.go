package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"qa_go/auth"
	"qa_go/cache"
	"qa_go/serializer"
)

// JwtRequired 需要在Header中传递token
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
		token, err := jwt.ParseWithClaims(userToken, &auth.JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtSecretKey, nil
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

		// 将Token放入Context, 用于注销添加黑名单
		c.Set("token", token.Raw)

		// 将token携带的用户id信息存入上下文
		if jwtStruct, ok := token.Claims.(*auth.JwtClaim); ok {
			c.Set("user_id", &jwtStruct.UserId)
		}
	}
}
