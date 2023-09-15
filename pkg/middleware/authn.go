package middleware

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/blog/pkg/jwt"
	"github.com/sjxiang/blog/pkg/serializer"
)

// ErrMissingHeader 表示 `Authorization` 请求头为空.
var ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")

// 认证
func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		userId, err := parseRequest(c)
		if err != nil {
			serializer.BuildResponse(c, err, nil)
			c.Abort()
			return 
		}

		c.Set("X-Username", userId)
		c.Next()
	}
}

func parseRequest(c *gin.Context) (string, error) {
    header := c.Request.Header.Get("Authorization")

    if len(header) == 0 {
        return "", ErrMissingHeader
    }

    var token string
    // 从请求头中取出 token
    fmt.Sscanf(header, "Bearer %s", &token)

    return jwt.ExtractAuth2Token(token)
}