package user

import (
	"context"
	"regexp"

	"github.com/jinzhu/copier"
	"github.com/sjxiang/blog/internal/blog/model"
	"github.com/sjxiang/blog/internal/blog/store"
	v1 "github.com/sjxiang/blog/pkg/api/blog/v1"
	"github.com/sjxiang/blog/pkg/errno"
)

// UserBiz 定义了 user 模块在 biz 层所实现的方法.
type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

// UserBizImpl 接口的实现.
type userBizImpl struct {
	ds store.IStore
}

// 确保 userBizImpl 实现了 UserBiz 接口.
var _ UserBiz = (*userBizImpl)(nil)

// New 创建一个实现了 UserBiz 接口的实例.
func New(ds store.IStore) *userBizImpl {
	return &userBizImpl{ds: ds}
}

// Create 是 UserBiz 接口中 `Create` 方法的实现.
func (b *userBizImpl) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)

	if err := b.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}

		return err
	}

	return nil
}