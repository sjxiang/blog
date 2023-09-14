package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	v1 "github.com/sjxiang/blog/pkg/api/blog/v1"
	"github.com/sjxiang/blog/pkg/errno"
	"github.com/sjxiang/blog/pkg/serializer"
	"github.com/sjxiang/blog/pkg/zop"
)

// ChangePassword 用来修改指定用户的密码
func (ctrl *UserController) ChangePassword(c *gin.Context) {
	
	zop.C(c).Infow("更改密码")
	
	var req v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		serializer.BuildResponse(c, errno.ErrBind, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(req); err != nil {
		serializer.BuildResponse(c, errno.ErrInvalidParameter.WithMessage(err.Error()), nil)
		return	
	}

	if err := ctrl.b.Users().ChangePassword(c, c.Param("name"), &req); err != nil {
		serializer.BuildResponse(c, err, nil)
		return
	}

	serializer.BuildResponse(c, nil, nil)
}