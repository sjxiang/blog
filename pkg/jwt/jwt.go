package jwt

import (
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth2Claims struct {
	Identity   string       `json:"identity"`
	
	jwt.RegisteredClaims
}

var (
	secretKey string    = "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5"
	once   sync.Once
)

func Init(secret string) {
	once.Do(func() {
		if secret != "" {
			secretKey = secret
		}
	})
}


// 生成
func GenerateAuth2Token(identity string) (string, error) {

	claims := &Auth2Claims{
		Identity:     identity,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "BLOG",                          // 签发人
			ExpiresAt: &jwt.NumericDate{       
				Time: time.Now().Add(time.Hour * 300),  // 过期时间
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签发
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}


// 提取
func ExtractAuth2Token(stateToken string) (identity string, err error) {
	authClaims := &Auth2Claims{}
	token, err := jwt.ParseWithClaims(stateToken, authClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Auth2Claims)
	if !(ok && token.Valid) {
		return "", err
	}

	return claims.Identity, nil
}



// // ParseRequest 从请求头中获取令牌，并将其传递给 Parse 函数以解析令牌.
// func ParseRequest(c *gin.Context) (string, error) {
// 	header := c.Request.Header.Get("Authorization")

// 	if len(header) == 0 {
// 		return "", ErrMissingHeader
// 	}

// 	var t string
// 	// 从请求头中取出 token
// 	fmt.Sscanf(header, "Bearer %s", &t)

// 	return Parse(t, config.key)
// }