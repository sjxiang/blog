package user

import (
	"github.com/sjxiang/blog/internal/blog/biz"
	"github.com/sjxiang/blog/internal/blog/store"
	"github.com/sjxiang/blog/pkg/auth"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	a *auth.Authz  // 授权器
    b biz.IBiz     // 
}

// New 创建一个 user controller.
func New(a *auth.Authz, ds store.IStore) *UserController {
    return &UserController{a: a, b: biz.NewBiz(ds)}
}

