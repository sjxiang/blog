package blog

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/blog/internal/blog/controller/v1/user"
	"github.com/sjxiang/blog/internal/blog/store"
	"github.com/sjxiang/blog/pkg/auth"
	"github.com/sjxiang/blog/pkg/errno"
	"github.com/sjxiang/blog/pkg/middleware"
	"github.com/sjxiang/blog/pkg/serializer"
	"github.com/sjxiang/blog/pkg/zop"
)


func setupRoute(authz *auth.Authz, store store.IStore, router *gin.Engine) error {
	
	// 构建 controller（依赖倒置）
	uc := user.New(authz, store)
	// pc
	

	// 注册业务路由
	router.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := router.Group("/v1")
	{	
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{	
			userv1.POST("", uc.Create)                               // 创建用户
			userv1.PUT("/:name/change-password", uc.ChangePassword)  // 修改用户密码
			
			userv1.Use(middleware.Authn(), middleware.Authz(authz))  // 牛批，仅对后面注册的路由起作用，横插一杠
		
			userv1.GET("/:name", uc.Get)        // 获取用户详情
			userv1.PUT("/:name", uc.Update)     // 更新用户
			userv1.GET("", uc.List)             // 列出用户列表，只有 root 用户才能访问
			userv1.DELETE("/:name", uc.Delete)  // 删除用户
		}
	}


	// 配置检查检查路由
	router.GET("/healthz", healthzHandler)

	// 配置 404 路由
	setupNoFoundHandler(router)

	return nil
}



func healthzHandler(ctx *gin.Context) {
	zop.C(ctx).Infow("健康检查被调用")

	serializer.BuildResponse(ctx, nil, map[string]string{"status": "ok"})
}

func setupNoFoundHandler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(ctx *gin.Context) {
		// 稍微沾点边，那就给点提示
		if strings.HasPrefix(ctx.Request.RequestURI, "/v1") || strings.HasPrefix(ctx.Request.RequestURI, "/api") {
			err := errno.ErrPageNotFound.WithMessage(
				fmt.Sprintf("路由未定义，请确认 url 和请求方法是否正确，Invalid URL (%s %s)", ctx.Request.Method, ctx.Request.URL.Path))
			serializer.BuildResponse(ctx, err, nil)
		}

		// 否则，给爷爬
	})
}
