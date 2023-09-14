package store

import (
	"context"

	"gorm.io/gorm"
	
	"github.com/sjxiang/blog/internal/blog/model"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
type UserStore interface {
    Create(ctx context.Context, user *model.UserM) error
}

// UserStore 接口的实现.
type userStoreImpl struct {
    db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ UserStore = (*userStoreImpl)(nil)

func newUsers(db *gorm.DB) *userStoreImpl {
    return &userStoreImpl{db}
}

// Create 插入一条 user 记录.
func (u *userStoreImpl) Create(ctx context.Context, user *model.UserM) error {
    return u.db.Create(&user).Error
}
