package middleware

import (
	"strings"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func CorsV1() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE", "PATCH", "OPTIONS"},
		
		// 在使用 JWT 的时候，因为我们使用了 Authorization 的头部，所以要加上
		AllowHeaders: []string{"Content-Type", "Authorization"},
		
		// 为了 JWT，你不加这个，前端是拿不到的
		ExposeHeaders: []string{"X-Jwt-Token"},
		
		// 是否允许你带 cookie 之类的东西
		AllowCredentials: true,
		
		// AllowOrigins: []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// 你的开发环境
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	})
}



func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Origin, Authorization")  
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return 
		}

		c.Next()
	}
}


/*

	Web 开发趋势：

		前后端分离，双方部署的 ip 地址不同，通过 nginx 转发请求
		
			前端 192.168.0.2
			后端 192.168.0.3


		但也有安全问题，CSRF 跨站请求伪造（用户在前端误操作，到其它第三方网站，被盗取 cookie 等，切记不可以瑟瑟）
		
		
		应对，浏览器同源策略

			浏览器会加塞个 options 请求，header 里有个 Origin 字段，询问下策略；
			那就回复它，有什么要求，让浏览器自己掂量。

	
	总结：处于安全考虑，浏览器的小动作
*/