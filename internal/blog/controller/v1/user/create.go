package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	v1 "github.com/sjxiang/blog/pkg/api/blog/v1"
	"github.com/sjxiang/blog/pkg/errno"
	"github.com/sjxiang/blog/pkg/serializer"
	"github.com/sjxiang/blog/pkg/zop"
)


// 创建一个新的用户
func (ctrl *UserController) Create(c *gin.Context) {
	zop.C(c).Infow("创建用户")
	
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

	if _, err := ctrl.a.AddNamedPolicy("p", req.Username, "/v1/users/"+req.Username, "(GET)|(POST)|(PUT)|(DELETE)"); err != nil {
		serializer.BuildResponse(c, err, nil)
        return
    }

	// 返回响应
	serializer.BuildResponse(c, nil, nil)
}


