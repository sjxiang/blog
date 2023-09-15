package store

import (
	"context"

	"gorm.io/gorm"
	
	"github.com/sjxiang/blog/internal/blog/model"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
type UserStore interface {
    Create(ctx context.Context, user *model.UserM) error
    Get(ctx context.Context, username string) (*model.UserM, error)
    Update(ctx context.Context, user *model.UserM) error
    List(ctx context.Context, username string, offset, limit int) (int64, []*model.PostM, error)
    Delete(ctx context.Context, username string, postIDs []string) error
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

// Get 根据用户名查询指定 user 的数据库记录.
func (u *userStoreImpl) Get(ctx context.Context, username string) (*model.UserM, error) {
    var user model.UserM                                                         
    if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err                                                            
    }                  
     
    return &user, nil
}                    
 
// Update 更新一条 user 数据库记录.
func (u *userStoreImpl) Update(ctx context.Context, user *model.UserM) error {
    return u.db.Save(user).Error
}

func (u *userStoreImpl) List(ctx context.Context, username string, offset, limit int) (int64, []*model.PostM, error) {
    panic("")
}

func (u *userStoreImpl) Delete(ctx context.Context, username string, postIDs []string) error {
    panic("")
}