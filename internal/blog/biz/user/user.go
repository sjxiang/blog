package user

import (
	"context"
	"regexp"

	"github.com/jinzhu/copier"
	"github.com/sjxiang/blog/internal/blog/model"
	"github.com/sjxiang/blog/internal/blog/store"
	v1 "github.com/sjxiang/blog/pkg/api/blog/v1"
	"github.com/sjxiang/blog/pkg/errno"
	"github.com/sjxiang/blog/pkg/jwt"
	"github.com/sjxiang/blog/pkg/util"
)

// UserBiz 定义了 user 模块在 biz 层所实现的方法.
type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
	ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	
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



// ChangePassword 是 UserBiz 接口中 `ChangePassword` 方法的实现.
func (b *userBizImpl) ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error {
    userM, err := b.ds.Users().Get(ctx, username)
    if err != nil {
        return err
    }

    if err := util.Compare(userM.Password, r.OldPassword); err != nil {
        return errno.ErrPasswordIncorrect
    }

    userM.Password, _ = util.Encrypt(r.NewPassword)
    if err := b.ds.Users().Update(ctx, userM); err != nil {
        return err
    }

    return nil
}

// Login 是 UserBiz 接口中 `Login` 方法的实现.
func (b *userBizImpl) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
    // 获取登录用户的所有信息
    user, err := b.ds.Users().Get(ctx, r.Username)
    if err != nil {
        return nil, errno.ErrUserNotFound
    }

    // 对比传入的明文密码和数据库中已加密过的密码是否匹配
    if err := util.Compare(user.Password, r.Password); err != nil {
        return nil, errno.ErrPasswordIncorrect
    }

    // 如果匹配成功，说明登录成功，签发 token 并返回
	t, err := jwt.GenerateAuth2Token(r.Username)
	if err != nil {
        return nil, errno.ErrSignToken
    }

    return &v1.LoginResponse{Token: t}, nil
}
