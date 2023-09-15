package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/blog/pkg/errno"
	"github.com/sjxiang/blog/pkg/serializer"
	"github.com/sjxiang/blog/pkg/zop"
)


// Auther 用来定义授权接口实现.
// sub: 操作主题，obj：操作对象, act：操作
type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

// Authz 是 Gin 中间件，用来进行请求授权.
func Authz(a Auther) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sub := ctx.GetString("X-Username")  // 访问实体，用户名
		obj := ctx.Request.URL.Path         // 访问资源，访问路径
		act := ctx.Request.Method           // 访问方法，HTTP 方法

		zop.Infow("构建授权上下文", "sub", sub, "obj", obj, "act", act)

		if allowed, _ := a.Authorize(sub, obj, act); !allowed {
			serializer.BuildResponse(ctx, errno.ErrUnauthorized, nil)
			ctx.Abort()
			return
		}

		// 有没有区别不大，执行到底，自动跳下一个 gin.HandlerFunc
		// ctx.Next()
	}
}

