package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/blog/internal/blog/biz"
	"github.com/sjxiang/blog/internal/blog/store"
	v1 "github.com/sjxiang/blog/pkg/api/blog/v1"
	"github.com/sjxiang/blog/pkg/errno"
	"github.com/sjxiang/blog/pkg/serializer"
	"github.com/sjxiang/blog/pkg/zop"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
    b biz.IBiz
}

// New 创建一个 user controller.
func New(ds store.IStore) *UserController {
    return &UserController{b: biz.NewBiz(ds)}
}

func (ctrl *UserController) Create(c *gin.Context) {
	zop.C(c).Infow("Create user called")
	
	// 参数解析
	var req v1.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		serializer.BuildResponse(c, errno.ErrBind, nil)
		return
    }
	// 参数校验
	if _, err := govalidator.ValidateStruct(req); err != nil {
		serializer.BuildResponse(c, errno.ErrInvalidParameter.WithMessage(err.Error()), nil)
        return
    }

	// 业务逻辑
	if err := ctrl.b.Users().Create(c, &req); err != nil {
		serializer.BuildResponse(c, err, nil)
		return 
	}

	// 返回响应
	serializer.BuildResponse(c, nil, nil)
}