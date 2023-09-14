package user

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/sjxiang/blog/pkg/api/blog/v1"
	"github.com/sjxiang/blog/pkg/errno"
	"github.com/sjxiang/blog/pkg/serializer"
	"github.com/sjxiang/blog/pkg/zop"
)


// 登录 blog 并返回一个 JWT Token.
func (ctrl *UserController) Login(c *gin.Context) {
	zop.C(c).Infow("登录")

	var req v1.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		serializer.BuildResponse(c, errno.ErrBind, nil)
		return 
	}

	resp, err := ctrl.b.Users().Login(c, &req)
	if err != nil {
		serializer.BuildResponse(c, err, nil)
	}

	serializer.BuildResponse(c, nil, resp)
}
