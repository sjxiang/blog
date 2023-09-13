package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 检查请求头中是否有 `X-Request-ID`，如果有则复用，没有则新建
		requestID := ctx.Request.Header.Get("X-Request-ID")

		if requestID == "" {
			requestID = uuid.New().String()
		}
		
        // 将 RequestID 保存在 gin.Context 中，方便后边程序使用                                                            
        ctx.Set("X-Request-ID", requestID)                                                                              

        // 将 RequestID 保存在 HTTP 返回头中，Header 的键为 `X-Request-ID`                                                                            
        ctx.Writer.Header().Set("X-Request-ID", requestID) 
		
		ctx.Next()
	}
}